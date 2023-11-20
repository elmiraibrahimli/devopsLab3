# devopsLab3


**Description**
This project involves setting up a web application with a separate PostgreSQL database backend. The infrastructure includes two AWS EC2 instances: one for the web application and another for the database. We use Ansible for automation, with two main roles: one to configure PostgreSQL on the database instance and another to deploy the web application in a Docker container on the app instance.

Infrastructure Setup
1. AWS EC2 Instances
Database Instance:
Used for hosting the PostgreSQL database.
Instance Details: 34.201.171.30

Application Instance:
Used for hosting the web application in a Docker container.
Instance Details: 54.89.153.168

Getting Started
Prerequisites
Ansible installed on your control machine.
SSH access to both AWS EC2 instances.
Docker installed on the Application instance.
Git (optional) to clone the repository containing the playbooks and roles.


The Inventory File
The inventory.ini file includes the IP addresses of my web and database servers:

[production]
54.89.153.168

[database]
34.201.171.30



Ansible Vault:
Created secrets.yml file which includes sensitive information such as database username and password and then encrypted it with vault:
ansible-vault create secrets.yml
**The password for vault is: vault**

You can read the content of secrets.yml file with decrypt function of vault:
ansible-vault decrypt secrets.yml







Running the Ansible Playbook

Navigate to the playbook directory.
Run the playbook:
I created one playbook (playbook.yml) and separate playbooks (database-playbook.yml and app-playbook.yml) for database and app. You can run the separate playbooks for only DB and App or run the playbook.yml for configuring these 2 instances at the same time.

ansible-playbook -i inventory.ini playbook.yml --ask-vault-pass
Enter the vault password when prompted (vault).

Accessing the Application
After the deployment, access the web application by navigating to http://54.89.153.168:8080/items/{id}  (http://54.89.153.168:8080/items/2) in your web browser.


Screenshots:
![1](https://github.com/elmiraibrahimli/devopsLab3/assets/94115234/b920e9fc-2ebb-433f-bc1d-18e6354098b9)



![2](https://github.com/elmiraibrahimli/devopsLab3/assets/94115234/61773a3d-4d44-4317-8fee-9cfab6ecd0c3)


Authors
Elmira Ibrahimli
