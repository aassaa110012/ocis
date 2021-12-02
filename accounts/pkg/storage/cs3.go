package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"

	"github.com/cs3org/reva/pkg/auth/scope"

	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	revactx "github.com/cs3org/reva/pkg/ctx"
	"github.com/cs3org/reva/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/pkg/token"
	"github.com/cs3org/reva/pkg/token/manager/jwt"
	"github.com/owncloud/ocis/accounts/pkg/config"
	"github.com/owncloud/ocis/accounts/pkg/proto/v0"
	olog "github.com/owncloud/ocis/ocis-pkg/log"
	"google.golang.org/grpc/metadata"
)

// CS3Repo provides a cs3 implementation of the Repo interface
type CS3Repo struct {
	cfg               *config.Config
	tm                token.Manager
	storageProvider   provider.ProviderAPIClient
	dataGatewayClient *http.Client
}

// NewCS3Repo creates a new cs3 repo
func NewCS3Repo(cfg *config.Config) (Repo, error) {
	tokenManager, err := jwt.New(map[string]interface{}{
		"secret": cfg.TokenManager.JWTSecret,
	})

	if err != nil {
		return nil, err
	}

	client, err := pool.GetStorageProviderServiceClient(cfg.Repo.CS3.ProviderAddr)
	if err != nil {
		return nil, err
	}

	return CS3Repo{
		cfg:               cfg,
		tm:                tokenManager,
		storageProvider:   client,
		dataGatewayClient: http.DefaultClient,
	}, nil
}

// WriteAccount writes an account via cs3 and modifies the provided account (e.g. with a generated id).
func (r CS3Repo) WriteAccount(ctx context.Context, a *proto.Account) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)
	if err := r.makeRootDirIfNotExist(ctx, accountsFolder); err != nil {
		return err
	}

	var by []byte
	if by, err = json.Marshal(a); err != nil {
		return err
	}

	err = r.uploadHelper(ctx, r.accountURL(a.Id), by)
	return err

}

// LoadAccount loads an account via cs3 by id and writes it to the provided account
func (r CS3Repo) LoadAccount(ctx context.Context, id string, a *proto.Account) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)

	return r.loadAccount(ctx, id, a)
}

// LoadAccounts loads all the accounts from the cs3 api
func (r CS3Repo) LoadAccounts(ctx context.Context, a *[]*proto.Account) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)
	res, err := r.storageProvider.ListContainer(ctx, &provider.ListContainerRequest{
		Ref: &provider.Reference{
			Path: path.Join("/meta", accountsFolder),
		},
	})
	if err != nil {
		return err
	}

	log := olog.NewLogger(olog.Pretty(r.cfg.Log.Pretty), olog.Color(r.cfg.Log.Color), olog.Level(r.cfg.Log.Level))
	for i := range res.Infos {
		acc := &proto.Account{}
		err := r.loadAccount(ctx, filepath.Base(res.Infos[i].Path), acc)
		if err != nil {
			log.Err(err).Msg("could not load account")
			continue
		}
		*a = append(*a, acc)
	}
	return nil
}

func (r CS3Repo) loadAccount(ctx context.Context, id string, a *proto.Account) error {
	account, err := r.downloadHelper(ctx, r.accountURL(id))
	if err != nil {
		switch err.(type) {
		case notFoundErr:
			return notFoundErr{"account", id}
		}
		return err
	}
	return json.Unmarshal(account, &a)
}

// DeleteAccount deletes an account via cs3 by id
func (r CS3Repo) DeleteAccount(ctx context.Context, id string) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)

	resp, err := r.storageProvider.Delete(ctx, &provider.DeleteRequest{
		Ref: &provider.Reference{
			Path: path.Join("/meta", accountsFolder, id),
		},
	})

	if err != nil {
		return err
	}

	// TODO Handle other error codes?
	if resp.Status.Code == v1beta11.Code_CODE_NOT_FOUND {
		return &notFoundErr{"account", id}
	}

	return nil
}

// WriteGroup writes a group via cs3 and modifies the provided group (e.g. with a generated id).
func (r CS3Repo) WriteGroup(ctx context.Context, g *proto.Group) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)
	if err := r.makeRootDirIfNotExist(ctx, groupsFolder); err != nil {
		return err
	}

	var by []byte
	if by, err = json.Marshal(g); err != nil {
		return err
	}

	err = r.uploadHelper(ctx, r.groupURL(g.Id), by)
	return err
}

// LoadGroup loads a group via cs3 by id and writes it to the provided group
func (r CS3Repo) LoadGroup(ctx context.Context, id string, g *proto.Group) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)

	return r.loadGroup(ctx, id, g)
}

// LoadGroups loads all the groups from the cs3 api
func (r CS3Repo) LoadGroups(ctx context.Context, g *[]*proto.Group) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)
	res, err := r.storageProvider.ListContainer(ctx, &provider.ListContainerRequest{
		Ref: &provider.Reference{
			Path: path.Join("/meta", groupsFolder),
		},
	})
	if err != nil {
		return err
	}

	log := olog.NewLogger(olog.Pretty(r.cfg.Log.Pretty), olog.Color(r.cfg.Log.Color), olog.Level(r.cfg.Log.Level))
	for i := range res.Infos {
		grp := &proto.Group{}
		err := r.loadGroup(ctx, filepath.Base(res.Infos[i].Path), grp)
		if err != nil {
			log.Err(err).Msg("could not load account")
			continue
		}
		*g = append(*g, grp)
	}
	return nil
}

func (r CS3Repo) loadGroup(ctx context.Context, id string, g *proto.Group) error {
	group, err := r.downloadHelper(ctx, r.groupURL(id))
	if err != nil {
		switch err.(type) {
		case notFoundErr:
			return notFoundErr{"group", id}
		}
		return err
	}
	return json.Unmarshal(group, &g)
}

// DeleteGroup deletes a group via cs3 by id
func (r CS3Repo) DeleteGroup(ctx context.Context, id string) (err error) {
	t, err := r.authenticate(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, revactx.TokenHeader, t)

	resp, err := r.storageProvider.Delete(ctx, &provider.DeleteRequest{
		Ref: &provider.Reference{
			Path: path.Join("/meta", groupsFolder, id),
		},
	})

	if err != nil {
		return err
	}

	// TODO Handle other error codes?
	if resp.Status.Code == v1beta11.Code_CODE_NOT_FOUND {
		return &notFoundErr{"group", id}
	}

	return err
}

func (r CS3Repo) authenticate(ctx context.Context) (token string, err error) {
	return AuthenticateCS3(ctx, r.cfg.ServiceUser, r.tm)
}

// AuthenticateCS3 mints an auth token for communicating with cs3 storage based on a service user from config
func AuthenticateCS3(ctx context.Context, su config.ServiceUser, tm token.Manager) (token string, err error) {
	u := &user.User{
		Id: &user.UserId{
			OpaqueId: su.UUID,
		},
		Groups:    []string{},
		UidNumber: su.UID,
		GidNumber: su.GID,
	}
	s, err := scope.AddOwnerScope(nil)
	if err != nil {
		return
	}
	return tm.MintToken(ctx, u, s)
}

func (r CS3Repo) accountURL(id string) string {
	return path.Join(accountsFolder, id)
}

func (r CS3Repo) groupURL(id string) string {
	return path.Join(groupsFolder, id)
}

func (r CS3Repo) makeRootDirIfNotExist(ctx context.Context, folder string) error {
	return MakeDirIfNotExist(ctx, r.storageProvider, folder)
}

// MakeDirIfNotExist will create a root node in the metadata storage. Requires an authenticated context.
func MakeDirIfNotExist(ctx context.Context, sp provider.ProviderAPIClient, folder string) error {
	var rootPathRef = &provider.Reference{
		Path: path.Join("/meta", folder),
	}

	resp, err := sp.Stat(ctx, &provider.StatRequest{
		Ref: rootPathRef,
	})

	if err != nil {
		return err
	}

	if resp.Status.Code == v1beta11.Code_CODE_NOT_FOUND {
		_, err := sp.CreateContainer(ctx, &provider.CreateContainerRequest{
			Ref: rootPathRef,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (r CS3Repo) uploadHelper(ctx context.Context, path string, content []byte) error {

	ref := provider.InitiateFileUploadRequest{
		Ref: &provider.Reference{
			Path: path,
		},
	}

	res, err := r.storageProvider.InitiateFileUpload(ctx, &ref)
	if err != nil {
		return err
	}

	var endpoint string

	for _, proto := range res.GetProtocols() {
		if proto.Protocol == "simple" {
			endpoint = proto.GetUploadEndpoint()
			break
		}
	}
	if endpoint == "" {
		return errors.New("metadata storage doesn't support the simple upload protocol")
	}

	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(content))
	if err != nil {
		return err
	}
	resp, err := r.dataGatewayClient.Do(req)
	if err != nil {
		return err
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

func (r CS3Repo) downloadHelper(ctx context.Context, path string) (content []byte, err error) {

	ref := provider.InitiateFileDownloadRequest{
		Ref: &provider.Reference{
			Path: path,
		},
	}

	res, err := r.storageProvider.InitiateFileDownload(ctx, &ref)
	if err != nil {
		return []byte{}, err
	}

	var endpoint string

	for _, proto := range res.GetProtocols() {
		if proto.Protocol == "simple" {
			endpoint = proto.GetDownloadEndpoint()
			break
		}
	}
	if endpoint == "" {
		return []byte{}, errors.New("metadata storage doesn't support the simple download protocol")
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := r.dataGatewayClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, &notFoundErr{}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if err = resp.Body.Close(); err != nil {
		return []byte{}, err
	}

	return b, nil
}
