testing with local ssh server
---------------------------
1) pull image
> docker pull rastasheep/ubuntu-sshd

2) Run the Container: Start a Docker container from the image.
sh

> docker run -d -p 2222:22 --name test-ssh-server rastasheep/ubuntu-sshd

This command starts a new container and forwards the host's port 2222 to the container's port 22, which is the default SSH port.

3) Connect via SSH: Use an SSH client to connect to the server.

> ssh root@localhost -p 2222
password is root
