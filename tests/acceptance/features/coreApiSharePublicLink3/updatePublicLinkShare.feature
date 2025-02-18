@api @files_sharing-app-required @public_link_share-feature-required @issue-ocis-reva-252
Feature: update a public link share

  Background:
    Given using OCS API version "1"
    And user "Alice" has been created with default attributes and without skeleton files


  Scenario Outline: API responds with a full set of parameters when owner changes the expireDate of a public share
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    When user "Alice" updates the last public link share using the sharing API with
      | expireDate | 2040-01-01T23:59:59+0100 |
    Then the OCS status code should be "<ocs_status_code>"
    And the OCS status message should be "OK"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                         | A_STRING             |
      | share_type                 | public_link          |
      | uid_owner                  | %username%           |
      | displayname_owner          | %displayname%        |
      | permissions                | read                 |
      | stime                      | A_NUMBER             |
      | parent                     |                      |
      | expiration                 | A_STRING             |
      | token                      | A_STRING             |
      | uid_file_owner             | %username%           |
      | displayname_file_owner     | %displayname%        |
      | additional_info_owner      | %emailaddress%       |
      | additional_info_file_owner | %emailaddress%       |
      | item_type                  | folder               |
      | item_source                | A_STRING             |
      | path                       | /FOLDER              |
      | mimetype                   | httpd/unix-directory |
      | storage_id                 | A_STRING             |
      | storage                    | A_NUMBER             |
      | file_source                | A_STRING             |
      | file_target                | /FOLDER              |
      | mail_send                  | 0                    |
      | name                       |                      |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |

  @smokeTest @issue-ocis-reva-336
  Scenario Outline: Creating a new public link share, updating its expiration date and getting its info
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has updated the last public link share with
      | expireDate | 2033-01-31T23:59:59+0100 |
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                | A_STRING             |
      | item_type         | folder               |
      | item_source       | A_STRING             |
      | share_type        | public_link          |
      | file_source       | A_STRING             |
      | file_target       | /FOLDER              |
      | permissions       | read                 |
      | stime             | A_NUMBER             |
      | expiration        | 2033-01-31           |
      | token             | A_TOKEN              |
      | storage           | A_STRING             |
      | mail_send         | 0                    |
      | uid_owner         | %username%           |
      | displayname_owner | %displayname%        |
      | url               | AN_URL               |
      | mimetype          | httpd/unix-directory |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: Creating a new public link share with password and adding an expiration date using public API
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has uploaded file with content "Random data" to "/randomfile.txt"
    And user "Alice" has created a public link share with settings
      | path     | randomfile.txt |
      | password | %public%       |
    When user "Alice" updates the last public link share using the sharing API with
      | expireDate | 2040-01-01T23:59:59+0100 |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the public should be able to download the last publicly shared file using the old public WebDAV API with password "%public%" and the content should be "Random data"
    And the public should be able to download the last publicly shared file using the new public WebDAV API with password "%public%" and the content should be "Random data"

    @issue-ocis-2079
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: Creating a new public link share with password and removing (updating) it to make the resources accessible without password using public API
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has uploaded file with content "Random data" to "/randomfile.txt"
    And user "Alice" has created a public link share with settings
      | path     | randomfile.txt |
      | password | %public%       |
    When user "Alice" updates the last public link share using the sharing API with
      #removing password is basically making password empty
      | password | %remove% |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the public should be able to download the last publicly shared file using the old public WebDAV API without a password and the content should be "Random data"
    And the public should be able to download the last publicly shared file using the new public WebDAV API without a password and the content should be "Random data"

    @issue-ocis-2079
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |

  @issue-ocis-reva-336
  Scenario Outline: Creating a new public link share, updating its password and getting its info
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has updated the last public link share with
      | password | %public% |
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                | A_STRING             |
      | item_type         | folder               |
      | item_source       | A_STRING             |
      | share_type        | public_link          |
      | file_source       | A_STRING             |
      | file_target       | /FOLDER              |
      | permissions       | read                 |
      | stime             | A_NUMBER             |
      | token             | A_TOKEN              |
      | storage           | A_STRING             |
      | mail_send         | 0                    |
      | uid_owner         | %username%           |
      | displayname_owner | %displayname%        |
      | url               | AN_URL               |
      | mimetype          | httpd/unix-directory |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |

  @issue-ocis-reva-336
  Scenario Outline: Creating a new public link share, updating its permissions and getting its info
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has updated the last public link share with
      | permissions | read,update,create,delete |
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                | A_STRING                  |
      | item_type         | folder                    |
      | item_source       | A_STRING                  |
      | share_type        | public_link               |
      | file_source       | A_STRING                  |
      | file_target       | /FOLDER                   |
      | permissions       | read,update,create,delete |
      | stime             | A_NUMBER                  |
      | token             | A_TOKEN                   |
      | storage           | A_STRING                  |
      | mail_send         | 0                         |
      | uid_owner         | %username%                |
      | displayname_owner | %displayname%             |
      | url               | AN_URL                    |
      | mimetype          | httpd/unix-directory      |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |

  @issue-ocis-reva-336
  Scenario Outline: Creating a new public link share, updating its permissions to view download and upload and getting its info
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has updated the last public link share with
      | permissions | read,update,create,delete |
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                | A_STRING                  |
      | item_type         | folder                    |
      | item_source       | A_STRING                  |
      | share_type        | public_link               |
      | file_source       | A_STRING                  |
      | file_target       | /FOLDER                   |
      | permissions       | read,update,create,delete |
      | stime             | A_NUMBER                  |
      | token             | A_TOKEN                   |
      | storage           | A_STRING                  |
      | mail_send         | 0                         |
      | uid_owner         | %username%                |
      | displayname_owner | %displayname%             |
      | url               | AN_URL                    |
      | mimetype          | httpd/unix-directory      |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |

  @issue-ocis-reva-336
  Scenario Outline: Creating a new public link share, updating publicUpload option and getting its info
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has updated the last public link share with
      | publicUpload | true |
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                | A_STRING                  |
      | item_type         | folder                    |
      | item_source       | A_STRING                  |
      | share_type        | public_link               |
      | file_source       | A_STRING                  |
      | file_target       | /FOLDER                   |
      | permissions       | read,update,create,delete |
      | stime             | A_NUMBER                  |
      | token             | A_TOKEN                   |
      | storage           | A_STRING                  |
      | mail_send         | 0                         |
      | uid_owner         | %username%                |
      | displayname_owner | %displayname%             |
      | url               | AN_URL                    |
      | mimetype          | httpd/unix-directory      |
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: Adding public upload to a read only shared folder as recipient is not allowed using the public API
    Given auto-accept shares has been disabled
    And using OCS API version "<ocs_api_version>"
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "/test"
    And user "Alice" has shared folder "/test" with user "Brian" with permissions "share,read"
    And user "Brian" has accepted share "/test" offered by user "Alice"
    And user "Brian" has created a public link share with settings
      | path         | /Shares/test |
      | publicUpload | false        |
    When user "Brian" updates the last public link share using the sharing API with
      | publicUpload | true |
    Then the OCS status code should be "404"
    And the HTTP status code should be "<http_status_code>"
    And uploading a file should not work using the old public WebDAV API
    And uploading a file should not work using the new public WebDAV API

    @issue-ocis-2079
    Examples:
      | ocs_api_version | http_status_code |
      | 1               | 200              |
      | 2               | 404              |


  Scenario Outline: Adding public upload to a shared folder as recipient is allowed with permissions using the public API
    Given auto-accept shares has been disabled
    And using OCS API version "<ocs_api_version>"
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "/test"
    And user "Alice" has shared folder "/test" with user "Brian" with permissions "all"
    And user "Brian" has accepted share "/test" offered by user "Alice"
    And user "Brian" has created a public link share with settings
      | path         | /Shares/test |
      | publicUpload | false        |
    When user "Brian" updates the last public link share using the sharing API with
      | publicUpload | true |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And uploading a file should work using the old public WebDAV API
    And uploading a file should work using the new public WebDAV API

    @issue-ocis-2079
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: Adding public link with all permissions to a read only shared folder as recipient is not allowed using the public API
    Given auto-accept shares has been disabled
    And using OCS API version "<ocs_api_version>"
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "/test"
    And user "Alice" has shared folder "/test" with user "Brian" with permissions "share,read"
    And user "Brian" has accepted share "/test" offered by user "Alice"
    And user "Brian" has created a public link share with settings
      | path        | /Shares/test |
      | permissions | read         |
    When user "Brian" updates the last public link share using the sharing API with
      | permissions | read,update,create,delete |
    Then the OCS status code should be "404"
    And the HTTP status code should be "<http_status_code>"
    And uploading a file should not work using the old public WebDAV API
    And uploading a file should not work using the new public WebDAV API

    @issue-ocis-2079
    Examples:
      | ocs_api_version | http_status_code |
      | 1               | 200              |
      | 2               | 404              |


  Scenario Outline: Adding public link with all permissions to a read only shared folder as recipient is allowed with permissions using the public API
    Given auto-accept shares has been disabled
    And using OCS API version "<ocs_api_version>"
    And user "Brian" has been created with default attributes and without skeleton files
    And user "Alice" has created folder "/test"
    And user "Alice" has shared folder "/test" with user "Brian" with permissions "all"
    And user "Brian" has accepted share "/test" offered by user "Alice"
    And user "Brian" has created a public link share with settings
      | path        | /Shares/test |
      | permissions | read         |
    When user "Brian" updates the last public link share using the sharing API with
      | permissions | read,update,create,delete |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And uploading a file should work using the old public WebDAV API
    And uploading a file should work using the new public WebDAV API

    @issue-ocis-2079
    Examples:
      | ocs_api_version | ocs_status_code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: Updating share permissions from change to read restricts public from deleting files using the public API
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "PARENT"
    And user "Alice" has created folder "PARENT/CHILD"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/PARENT/CHILD/child.txt"
    And user "Alice" has created a public link share with settings
      | path        | /PARENT                   |
      | permissions | read,update,create,delete |
    And user "Alice" has updated the last public link share with
      | permissions | read |
    When the public deletes file "CHILD/child.txt" from the last public link share using the old public WebDAV API
    And the public deletes file "CHILD/child.txt" from the last public link share using the new public WebDAV API
    And the HTTP status code of responses on all endpoints should be "403"
    And as "Alice" file "PARENT/CHILD/child.txt" should exist

    @issue-ocis-2079 @issue-ocis-reva-292
    Examples:
      | ocs_api_version |
      | 1               |
      | 2               |


  Scenario Outline: Updating share permissions from read to change allows public to delete files using the public API
    Given using OCS API version "<ocs_api_version>"
    And user "Alice" has created folder "PARENT"
    And user "Alice" has created folder "PARENT/CHILD"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/PARENT/parent.txt"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/PARENT/CHILD/child.txt"
    And user "Alice" has created a public link share with settings
      | path        | /PARENT |
      | permissions | read    |
    And user "Alice" has updated the last public link share with
      | permissions | read,update,create,delete |
    When the public deletes file "CHILD/child.txt" from the last public link share using the new public WebDAV API
    And the public deletes file "parent.txt" from the last public link share using the new public WebDAV API
    Then the HTTP status code of responses on all endpoints should be "204"
    And as "Alice" file "PARENT/CHILD/child.txt" should not exist
    And as "Alice" file "PARENT/parent.txt" should not exist
    Examples:
      | ocs_api_version |
      | 1               |
      | 2               |


  Scenario Outline: API responds with a full set of parameters when owner renames the folder with a public link in ocis
    Given using OCS API version "<ocs_api_version>"
    And using <dav-path> DAV path
    And user "Alice" has created folder "FOLDER"
    And user "Alice" has created a public link share with settings
      | path | FOLDER |
    And user "Alice" has moved folder "/FOLDER" to "/RENAMED_FOLDER"
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                     | A_STRING             |
      | share_type             | public_link          |
      | uid_owner              | %username%           |
      | displayname_owner      | %displayname%        |
      | permissions            | read                 |
      | stime                  | A_NUMBER             |
      | parent                 |                      |
      | expiration             |                      |
      | token                  | A_STRING             |
      | uid_file_owner         | %username%           |
      | displayname_file_owner | %displayname%        |
      | item_type              | folder               |
      | item_source            | A_STRING             |
      | path                   | /RENAMED_FOLDER      |
      | mimetype               | httpd/unix-directory |
      | storage_id             | A_STRING             |
      | storage                | A_STRING             |
      | file_source            | A_STRING             |
      | file_target            | /RENAMED_FOLDER      |
      | mail_send              | 0                    |
      | name                   |                      |
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | old      | 1               | 100             |
      | old      | 2               | 200             |
      | new      | 1               | 100             |
      | new      | 2               | 200             |

    @personalSpace
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | spaces   | 1               | 100             |
      | spaces   | 2               | 200             |


  Scenario Outline: API responds with a full set of parameters when owner renames the file with a public link in ocis
    Given using OCS API version "<ocs_api_version>"
    And using <dav-path> DAV path
    And user "Alice" has uploaded file with content "some content" to "/lorem.txt"
    And user "Alice" has created a public link share with settings
      | path | lorem.txt |
    And user "Alice" has moved file "/lorem.txt" to "/new-lorem.txt"
    When user "Alice" gets the info of the last public link share using the sharing API
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    And the fields of the last response to user "Alice" should include
      | id                     | A_STRING       |
      | share_type             | public_link    |
      | uid_owner              | %username%     |
      | displayname_owner      | %displayname%  |
      | permissions            | read           |
      | stime                  | A_NUMBER       |
      | parent                 |                |
      | expiration             |                |
      | token                  | A_STRING       |
      | uid_file_owner         | %username%     |
      | displayname_file_owner | %displayname%  |
      | item_type              | file           |
      | item_source            | A_STRING       |
      | path                   | /new-lorem.txt |
      | mimetype               | text/plain     |
      | storage_id             | A_STRING       |
      | storage                | A_STRING       |
      | file_source            | A_STRING       |
      | file_target            | /new-lorem.txt |
      | mail_send              | 0              |
      | name                   |                |
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | old      | 1               | 100             |
      | old      | 2               | 200             |
      | new      | 1               | 100             |
      | new      | 2               | 200             |

    @personalSpace
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | spaces   | 1               | 100             |
      | spaces   | 2               | 200             |


  Scenario Outline: Updating the role of a public link to internal
    Given using OCS API version "<ocs_api_version>"
    And using <dav-path> DAV path
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/textfile.txt"
    And user "Alice" has created a public link share with settings
      | path        | /textfile.txt |
      | permissions | read          |
    When user "Alice" updates the last public link share using the sharing API with
      | permissions | 0 |
    Then the OCS status code should be "<ocs_status_code>"
    And the HTTP status code should be "200"
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | old      | 1               | 100             |
      | old      | 2               | 200             |
      | new      | 1               | 100             |
      | new      | 2               | 200             |

    @personalSpace
    Examples:
      | dav-path | ocs_api_version | ocs_status_code |
      | spaces   | 1               | 100             |
      | spaces   | 2               | 200             |
