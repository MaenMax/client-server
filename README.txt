Author: Maen Abu Hammour April,29,2022

This Golang client/server module comes as an implementation of Bloxroute client-server assignment.


Requirements:

 - Golang 1.16
 - Locally running NATS cluser on ports: 4222, 5222, and 6222

Compilation:
 - Run the Makefile using Linux command:  $ make
 from the home directory of the git repo

- To deploy the server (create tar.bz2 archive), run $ make deploy

- To clear generated binaries/archives from project, run $ make clean


Running the server:

From the home directory, execute:
$ ./bin/server


To add an item to the queue, you will need a publisher, which is an independent Go program used to publish the message/item to the locally running NATS server in order to be consumed and processed by the server. The source code of the publisher exists in "utils" folder


The client/publisher needs to provide its id, request_type, item id, and the data to be stored in the queue (string format)

Request Types:

1:AddItem  2:RemoveItem 3:GetItem 4:GetAllItems


Example of adding an item:
./bin/publisher -msg '{"client_id":"1","item": {"id": "1", "data": "My first item"},"request_type":1 }':

Example of removing an item:
./bin/publisher -msg '{"client_id":"1","item": {"id": "1"},"request_type":2}':

Example of getting an item:
./bin/publisher -msg '{"client_id":"1","item": {"id": "1"},"request_type":3}':

Example of getting all items:
./bin/publisher -msg '{"client_id":"1","item": {"id": "1"},"request_type":4}':


NOTE: Generated logs exist in log/server.log file.
