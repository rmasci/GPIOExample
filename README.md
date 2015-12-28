# GPIOExample
First off, most of this program was written by Dave Cheney, I pulled it out of his gpio/examples/blink.go. Thanks Dave!!

Example of GPIO Programming on Raspberry PI.  When you run this you should have your terminal screen and the WebIOPi web page from
your Raspberry PI displayed.  This is the blink lights example in a web browser. No soldering, resistors... What this shows is how potentially easy it is to program a Raspberry PI using GO.  Imagine each one of these GPIO Ports are connected to
a relay, and those relays could then operate light switches, garage doors, toasters, etc. From the looks of it on the web, connecting a 
relay doesn't look to be very difficult, and you can buy jumper wires to make the project solderless.

<h3>Install Go on Raspberry PI</h3>
Easiest way right now is to install the Go 1.6beta (Don't worry about the beta, for what we're doing, it's going to be stable). 
Just download and extract to /usr/local:
<pre>
wget https://storage.googleapis.com/golang/go1.6beta1.linux-arm6.tar.gz
sudo su -
cd /usr/local
sudo tar zxvf /home/pi/go1.6beta1.linux-arm6.tar.gz 
cd 
</pre>
## Setup Home Environment
Next I put some entries in the ~/.bashrc:
<pre>
export GOROOT=/usr/local/go
export GOBIN=/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=/usr/local/bin:~/bin:$PATH:$GOBIN
</pre>

Once you have that, either log in/out or type:
<pre>
source ~/.bashrc
</pre>

Make a go directory in your home, and install Dave Cheney's gpio:
<pre>
mkdir ~/go
go get github.com/davecheney/gpio
go get github.com/davecheney/gpio/rpi
</pre>

## Install WebIOPi:
http://webiopi.trouch.com/INSTALL.html

## Get GPIOExample
<pre>
wget https://github.com/rmasci/GPIOExample/archive/master.zip
unzip master.zip
cd GPIOExample-master
go build GPIOExample.go
</pre>

## To run
First open a web browser to http://<your PI's ip address>:8000/app/gpio-header
You'll see something that looks like the GPIO Pins on your Raspberry PI, next to some of the pins will be an IN.  When 
GPIOExample is run, you'll see these change from IN to OUT, and then back to IN.
In the terminal:
<pre>
sudo ./GPIOExample
</pre>

