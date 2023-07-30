<!-- Copyright (c) 2021 Dell Inc. or its subsidiaries. All Rights Reserved. -->
# srs-mock-service 
Mock Service for the SRS

Steps to generate server code:
1. Edit the openapi specification which contains the API definitions.
For that, edit the file ./swagger-ui/openapi.json 

2. [OPTIONAL] To validate the openapi specification:
make validate

3. To generate the server code for API interfaces, models, handlers and service functions:
make generate

4. To install dependencies and setup the new repo, and compile the code:
make install

5. [OPTIONAL] To re-compile code if edited:
make compile

4. To start the server:
make start

5. To stop the server:
make stop

6. To generate code, install dependencies, compile the code and start the server in one shot:
make all
 
7. To add implementation to service functions:
    a. Edit the files src/api/api_<tag>_service.go to add the API implementation.
    b. To prevent code generation from overwriting these files, include these files in the .openapi-generator-ignore (similar to .gitignore)
    c. Commit these files to the git so that your implementation is preserved and never overwritten by code generation

# File-Server

1. File server is implemented using https://github.com/codeskyblue/gohttpserver gohttps server, where we can upload and delete files under /mnt/srs-mock-http/upload.

2. If fileserver is not running on srs vm, you can start it using the script 'setup_fileserver.sh'
   $ sudo ./setup_fileserver.sh

3. File server is running at http://srs-mock-01.cec.lab.emc.com:80

4. To  upload, list, get file/files use following curl commands:

   Upload a file from srs-mock VM:
   upload a file named foo.txt to directory somedir
   curl -F file=@foo.txt  localhost:80/upload/somedir

   List files under folder:
   curl -v http://localhost:80/path-to-directory?json=true

   Get file content:
   curl -v http://localhost:80/path-to-file

5. For accessing fileserver from any vm other than srs-mock vm, then you can access it at http://srs-mock-01.cec.lab.emc.com:80

6. File older than 2 weeks will be removed from fileserver. File purging policy is applied using logrotate.
   You can see file /etc/logrotate.d/fileserver on srs-mock vm for file purging policy.

7. Custom script for applying file purging policy.

/mnt/srs-mock-http/upload/*.log
{
        daily
        rotate 0
        firstaction
          /usr/bin/find /mnt/srs-mock-http/upload/ -name "*.log" -mtime +14 -delete
        endscript
        nocreate
        missingok
        notifempty
}
Note: This custom script is subject to change


Run `./dev-image-deployment.sh` to deploy the service as a docker container in the VM
