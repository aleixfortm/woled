<p align="center">
  <img src="https://github.com/aleixfortm/gowol/assets/95043218/1bce0a50-3bdb-42ef-82c5-ce158e18938f" alt="storymous_tree" width="200" height="auto">
</p>



### **<p align="center">WOLed</p>**


<p align="center"><strong>> An over-engineered CLI Wake-on-LAN tool</strong><br>Create own device list and broadcast WOL packages through a CLI application</p>

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" /
  </a>
</div>


<br>

## Usage
```go
woled <command>
```

## Commands and flags
```go
  wol         Broadcast WOL packet to local network
  config      Show and tweak default app settings, including UDP Port and broadcast IP address
  list        Display a list of saved devices
  save        Save device data to local storage
  remove      Remove saved device
```
```go
  -h, --help   show arguments and examples
```

<br>

## <code>wol</code> [name]
### <strong>Description</strong>
Send a WOL packet to the local network, broadcasted to IP 255.255.255.255 by default.

### <strong>Usage</strong>
```python
  woled wol [name]
```
### <strong>Arguments</strong>
```go
  [name]  string  Name of the device
```
### <strong>Examples</strong>
```go
  woled wol PC-1
  woled wol "My computer"
```

<br>

## <code>list</code>
### <strong>Description</strong>
Display a list of previously saved devices from local storage file data.json

### <strong>Usage</strong>
```python
  woled list
```
### <strong>Arguments</strong>
```go
  None
```
### <strong>Examples</strong>
```go
  woled list
```

<br>

## <code>save</code> [name] [MAC]
### <strong>Description</strong>
Save your device to a local data file by specifying a name and the MAC address of the device.

### <strong>Usage</strong>
```python
  woled save [name] [MAC]
```
### <strong>Arguments</strong>
```go
  [name]  string   Name of the device
  [MAC]   string   MAC address of the device
```
### <strong>Examples</strong>
```go
  gowol save PC-1 00:11:22:33:44:55
  gowol save "My computer" 04:AF:32:4B:4C:95
```

<br>

## <code>remove</code> [name]
### <strong>Description</strong>
Send a WOL packet to the local network, broadcasted to IP 255.255.255.255 by default.

### <strong>Usage</strong>
```python
  woled wol [name]
```
### <strong>Arguments</strong>
```go
  [name]  string  Name of the device
```
### <strong>Examples</strong>
```go
  woled wol PC-1
  woled wol "My computer"
```

<br>
