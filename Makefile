install:
	sudo mv virsh-api-server.service /etc/systemd/system/virsh-api-server.service
	sudo wget -O /Virsh-Api-Server/Virsh-Api-Server https://github.com/yoanndelattre/Virsh-Api-Server/releases/download/v1/Virsh-Api-Server
	sudo chmod u+x /Virsh-Api-Server/Virsh-Api-Server
	sudo chmod u+x /Virsh-Api-Server/getVmStatus.sh
	sudo chmod u+x /Virsh-Api-Server/start.sh
	sudo systemctl enable virsh-api-server
	sudo systemctl start virsh-api-server