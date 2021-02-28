# JSON CLI Search

----
A simple CLI search on JSON files for User, Organisation and Tickets.
----

####
The program uses a cli library for command-line implementation - github.com/mitchellh/cli v1.1.2
The search is performed using GO reflection library to search by struct field names.
The dependent libraries are vendored and checked into the repository.
Search keys are case insensitive also underscores are ignored, e.g id, ID, _id are treated the same.


#### Steps to run the application:

---

###### _Prerequisites_
- GO go1.14.11
- Docker
---

###### Steps
- 1 Build
    ```bash
    go build search.go
    ```

- 2 Help command
```bash
./search help

    search-org          Search Organisation using the search key and field. Key and value are case and underscore agnostic
                        Syntax search-org <key> <value>
                        e.g: search-org id 1
    search-tkt          Search Ticket and it assigned user and organisation using the search key and field. Key and value are case and underscore agnostic
                        Syntax search-tkt <key> <value>
                        e.g: search-tkt id 1
    search-user         Search User and its organisation using the search key and field. Key and value are case and underscore agnostic
                        Syntax search-user <key> <value>
                        e.g: search-tkt id 1

```
- 3 Run search commands
```bash
./search <search-cmd> <key> <value>
```

#### Run Tests

```bash
make test
```
___

#### Docker Run
1. Builder docker image
   docker build -t search-app .
2. Run help
   ```bash
   docker run search-app ./search help
   ```
3. Search entity
   ```bash
      docker run search-app ./search <command> <key> <value>
   ```
   Example:
   ```bash
         docker run search-app ./search search-user id 1
         docker run search-app ./search search-org id 101
         docker run search-app ./search search-tkt id xxx-xxxx-xxx
      ```

#### Examples

###### search-user <key> <value>
````bash
json-search-cli:>./search search-user _id 1
[ {
  "_id" : "1",
  "shared" : "false",
  "last_login_at" : "2013-08-04T01:03:27 -10:00",
  "role" : "admin",
  "signature" : "Don't Worry Be Happy!",
  "timezone" : "Sri Lanka",
  "verified" : "true",
  "created_at" : "2016-04-15T05:19:46 -10:00",
  "active" : "true",
  "external_id" : "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "locale" : "en-AU",
  "url" : "http://xxxxxx.com/api/v2/users/1.json",
  "suspended" : "true",
  "tags" : [ "Springville", "Sutton", "Hartsville/Hartley", "Diaperville" ],
  "phone" : "8335-422-718",
  "organisationObj" : {
    "_id" : "119",
    "shared_tickets" : "false",
    "name" : "Multron",
    "created_at" : "2016-02-29T03:45:12 -11:00",
    "external_id" : "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "details" : "Non profit",
    "url" : "http://xxxxxx.com/api/v2/organizations/119.json",
    "domain_names" : [ "bleeko.com", "pulze.com", "xoggle.com", "sultraxin.com" ],
    "tags" : [ "Erickson", "Mccoy", "Wiggins", "Brooks" ]
  },
  "name" : "xxxxxx xxx",
  "alias" : "Miss xxxxxx",
  "email" : "xxxxxxx@xxxxx.com"
} ]
````

- ###### search-org <key> <value>
````bash
json-search-cli:>./search search-org _id 101
[ {
  "_id" : "101",
  "shared_tickets" : "false",
  "name" : "XXXXXXX",
  "created_at" : "2016-05-21T11:10:28 -10:00",
  "external_id" : "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "details" : "XXXXXXX",
  "url" : "http://xxxxxxxx.com/api/v2/organizations/101.json",
  "domain_names" : [ "kage.com", "ecratic.com", "endipin.com", "zentix.com" ],
  "tags" : [ "Fulton", "West", "Rodriguez", "Farley" ]
} ]
````

- ###### search-ticket <key> <value>
````bash
json-search-cli:>./search ticket-search _id xxx-xxxx-xxx
[ {
      "_id": "xxx-xxxx-xxx",
      "url": "http://xx.xxx.com",
      "external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
      "created_at": "2016-04-28T11:19:34 -10:00",
      "type": "incident",
      "subject": "A Catastrophe in Korea (North)",
      "description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
      "priority": "high",
      "status": "pending",
      "submitter_id": 38,
      "assignee_id": 24,
      "organization_id": 116,
      "tags": [
        "Ohio",
        "Pennsylvania",
        "American Samoa",
        "Northern Mariana Islands"
      ],
      "has_incidents": false,
      "due_at": "2016-07-31T02:37:50 -10:00",
      "via": "web"
    }
]
````

#### Scope of further improvements

- Take JSON files as input to the application
- making use of a map data structure for faster search
- Add tabular output format
