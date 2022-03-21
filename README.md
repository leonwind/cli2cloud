# <a href="https://cli2cloud.com">Cli2Cloud</a>


<p align="center">
<a href="https://cli2cloud.com">
    <img src="webapp/src/assets/cloudWhite.png" width="128" alt="Cli2Cloud Logo"/>
</a>
<br/>
Monitor and Share Your Terminal Output with Everyone from Everywhere in Realtime.
</p>

## Installation
Install the terminal client directly from the source by running
```bash
go install github.com/leonwind/cli2cloud/cli/cli2cloud@latest
```

## Examples
### Normal usage
To just pipe your terminal output from any arbitrary command, run
```bash
$ ping google.com | cli2cloud
Your client ID: 4sYe3G
Share and monitor it live from https://cli2cloud.com/4sYe3G

PING google.com (172.217.22.142): 56 data bytes
64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=12.306 ms
64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=14.317 ms
...
```

and open `https://cli2cloud.com/{your ID}` on any browser you have.
It will pipe both your `Stdout` and your `Stderr` output to the web.

### End-to-End encryption
Use the `-encrypt {password}` option to encrypt your data End-to-End using the AES CBC Mode and a 256 bit key generated based on your password using the PBKDF2 function.
```bash
$ ping google.com | cli2cloud -encrypt 1234
Your client ID: CGYWdD
Share and monitor it live from https://cli2cloud.com/CGYWdD#key=1234

PING google.com (172.217.22.142): 56 data bytes
64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=14.154 ms
64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=12.565 ms
...
```

To decrypt the data on the web, you need to enter the same password again. The server does not store your password or the hash of it and thus can't validate if your password is either correct or incorrect. You will see complete garbage if you enter a wrong password :)

Use the option `-encrypt-random` to generate a random secure password with 16 symbols.
```bash
$ ping google.com | cli2cloud -encrypt-random
Your password: mruI3ubFXTww1QYf
Your client ID: 56xY35
Share and monitor it live from https://cli2cloud.com/56xY35#key=mruI3ubFXTww1QYf

PING google.com (142.250.201.174): 56 data bytes
64 bytes from 142.250.201.174: icmp_seq=0 ttl=116 time=3.322 ms
64 bytes from 142.250.201.174: icmp_seq=1 ttl=116 time=2.648 ms
..mruI3ubFXTww1QYf.
```

## Feedback
Feel free to open a new [Issue](https://github.com/leonwind/cli2cloud/issues) regarding any feedback, bugs or feature requests.
