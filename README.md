# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.



## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)



PI Configuration: 
FILE : ".config/autostart/ChromiumPanel.desktop"
Notes: By running it in kiosk mode we will not get any notifications on brower

[Desktop Entry]
Version=1.0
Type=Application
Name=ChromiumPanel
Exec=/usr/bin/chromium-browser --noerrdialogs --kiosk --app=http://localhost:3000 --incognito


FILE : "/etc/rc.local"
Notes: Start the web app in deamon mode as soon as the PI boot up 

/home/pi/start_zephyr.sh 


FILE : ".config/lxsession/LXDE-pi/autostart"

@lxpanel --profile LXDE-pi
@pcmanfm --desktop --profile LXDE-pi
@xscreensaver -no-splash
@point-rpi
@unclutter -idle 0

FILE: "/boot/cmdline.txt"

dwc_otg.lpm_enable=0 console=serial0,115200 console=tty1 root=PARTUUID=13d159a1-02 rootfstype=ext4 elevator=deadline fsck.repair=yes rootwait quiet splash plymouth.ignore-serial-consoles
logo.nologo
consoleblank=0 loglevel=0 quiet
