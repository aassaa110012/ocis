# Userlog Service

The `userlog` service is a mediator between the `eventhistory` service and clients who want to be informed about user related events. It provides an API to retrieve those.

## Prerequisites

Running the `userlog` service without running the `eventhistory` service is not possible.

## Storing

The `userlog` service persists information via the configured store in `USERLOG_STORE_TYPE`. Possible stores are:
  -   `mem`: Basic in-memory store and the default.
  -   `ocmem`: Advanced in-memory store allowing max size.
  -   `redis`: Stores data in a configured redis cluster.
  -   `etcd`: Stores data in a configured etcd cluster.
  -   `nats-js`: Stores data using key-value-store feature of [nats jetstream](https://docs.nats.io/nats-concepts/jetstream/key-value-store)
  -   `noop`: Stores nothing. Useful for testing. Not recommended in productive enviroments.

1.  Note that in-memory stores are by nature not reboot persistent.
2.  Though usually not necessary, a database name and a database table can be configured for event stores if the event store supports this. Generally not applicapable for stores of type `in-memory`. These settings are blank by default which means that the standard settings of the configured store applies.
3.  The userlog service can be scaled if not using `in-memory` stores and the stores are configured identically over all instances.

## Configuring

For the time being, the configuration which user related events are of interest is hardcoded and cannot be changed.

## Retrieving

The `userlog` service provides an API to retrieve configured events. For now, this API is mostly following the [oc10 notification GET API](https://doc.owncloud.com/server/next/developer_manual/core/apis/ocs-notification-endpoint-v1.html#get-user-notifications).

## Deleting

To delete events for an user, use a `DELETE` request to `ocs/v2.php/apps/notifications/api/v1/notifications` containing the IDs to delete.
