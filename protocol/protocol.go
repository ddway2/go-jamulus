/*
Protocol message definition
---------------------------

- All messages received need to be acknowledged by an acknowledge packet (except
  of connection less messages)



MAIN FRAME
----------

    +-------------+------------+------------+------------------+ ...
    | 2 bytes TAG | 2 bytes ID | 1 byte cnt | 2 bytes length n | ...
    +-------------+------------+------------+------------------+ ...
        ... --------------+-------------+
        ...  n bytes data | 2 bytes CRC |
        ... --------------+-------------+

- TAG is an all zero bit word to identify protocol messages
- message ID defined by the defines PROTMESSID_x
- cnt: counter which is increment for each message and wraps around at 255
- length n in bytes of the data
- actual data, dependent on message type
- 16 bits CRC, calculated over the entire message and is transmitted inverted
  Generator polynom: G_16(x) = x^16 + x^12 + x^5 + 1, initial state: all ones



SPLIT MESSAGE CONTAINER
-----------------------

    +------------+------------------------+------------------+--------------+
    | 2 bytes ID | 1 byte number of parts | 1 byte split cnt | n bytes data |
    +------------+------------------------+------------------+--------------+

- ID is the message ID of the message being split
- number of parts - total number of parts comprising the whole message
- split cnt - number within number total for this part of the message
- data - subset of the data part of the original message being split



MESSAGES (with connection)
--------------------------

- PROTMESSID_ACKN: Acknowledgement message

    +------------------------------------------+
    | 2 bytes ID of message to be acknowledged |
    +------------------------------------------+

    note: the cnt value is the same as of the message to be acknowledged


- PROTMESSID_JITT_BUF_SIZE: Jitter buffer size

    +--------------------------+
    | 2 bytes number of blocks |
    +--------------------------+


- PROTMESSID_REQ_JITT_BUF_SIZE: Request jitter buffer size

    note: does not have any data -> n = 0


- PROTMESSID_CLIENT_ID: Sends the current client ID to the client

    +---------------------------------+
    | 1 byte channel ID of the client |
    +---------------------------------+


- PROTMESSID_CHANNEL_GAIN: Gain of channel

    +-------------------+--------------+
    | 1 byte channel ID | 2 bytes gain |
    +-------------------+--------------+


- PROTMESSID_CHANNEL_PAN: Gain of channel

    +-------------------+-----------------+
    | 1 byte channel ID | 2 bytes panning |
    +-------------------+-----------------+


- PROTMESSID_MUTE_STATE_CHANGED: Mute state of your signal at another client has changed

    +-------------------+-----------------+
    | 1 byte channel ID | 1 byte is muted |
    +-------------------+-----------------+


- PROTMESSID_CONN_CLIENTS_LIST: Information about connected clients

    for each connected client append following data:

    +-------------------+-----------------+--------------------+ ...
    | 1 byte channel ID | 2 bytes country | 4 bytes instrument | ...
    +-------------------+-----------------+--------------------+ ...
        ... --------------------+--------------------+ ...
        ...  1 byte skill level | 4 bytes IP address | ...
        ... --------------------+--------------------+ ...
        ... ------------------+---------------------------+
        ...  2 bytes number n | n bytes UTF-8 string name |
        ... ------------------+---------------------------+
        ... ------------------+---------------------------+
        ...  2 bytes number n | n bytes UTF-8 string city |
        ... ------------------+---------------------------+


- PROTMESSID_REQ_CONN_CLIENTS_LIST: Request connected clients list

    note: does not have any data -> n = 0


- PROTMESSID_CHANNEL_INFOS: Information about the channel

    +-----------------+--------------------+ ...
    | 2 bytes country | 4 bytes instrument | ...
    +-----------------+--------------------+ ...
        ... --------------------+ ...
        ...  1 byte skill level | ...
        ... --------------------+ ...
        ... ------------------+---------------------------+ ...
        ...  2 bytes number n | n bytes UTF-8 string name | ...
        ... ------------------+---------------------------+ ...
        ... ------------------+---------------------------+
        ...  2 bytes number n | n bytes UTF-8 string city |
        ... ------------------+---------------------------+


- PROTMESSID_REQ_CHANNEL_INFOS: Request infos of the channel

    note: does not have any data -> n = 0


- PROTMESSID_CHAT_TEXT: Chat text

    +------------------+----------------------+
    | 2 bytes number n | n bytes UTF-8 string |
    +------------------+----------------------+


- PROTMESSID_NETW_TRANSPORT_PROPS: Properties for network transport

    +------------------------+-------------------------+-----------------+ ...
    | 4 bytes base netw size | 2 bytes block size fact | 1 byte num chan | ...
    +------------------------+-------------------------+-----------------+ ...
        ... ------------------+-----------------------+ ...
        ...  4 bytes sam rate | 2 bytes audiocod type | ...
        ... ------------------+-----------------------+ ...
        ... ---------------+----------------------+
        ...  2 bytes flags | 4 bytes audiocod arg |
        ... ---------------+----------------------+

    - "base netw size":  length of the base network packet (frame) in bytes
    - "block size fact": block size factor
    - "num chan":        number of channels of the audio signal, e.g. "2" is
                         stereo
    - "sam rate":        sample rate of the audio stream
    - "audiocod type":   audio coding type, the following types are supported:
                          - 0: none, no audio coding applied
                          - 1: CELT
                          - 2: OPUS
                          - 3: OPUS64
    - "flags":           flags indicating network properties:
                          - 0: none
                          - 1: WITH_COUNTER (a packet counter is added to the audio packet)
    - "audiocod arg":    argument for the audio coder, if not used this value
                         shall be set to 0


- PROTMESSID_REQ_NETW_TRANSPORT_PROPS: Request properties for network transport

    note: does not have any data -> n = 0


- PROTMESSID_REQ_SPLIT_MESS_SUPPORT: Request split message support

    note: does not have any data -> n = 0


- PROTMESSID_SPLIT_MESS_SUPPORTED: Split messages are supported

    note: does not have any data -> n = 0


- PROTMESSID_LICENCE_REQUIRED: Licence required to connect to the server

    +---------------------+
    | 1 byte licence type |
    +---------------------+


- PROTMESSID_VERSION_AND_OS: Version number and operating system

    +-------------------------+------------------+------------------------------+
    | 1 byte operating system | 2 bytes number n | n bytes UTF-8 string version |
    +-------------------------+------------------+------------------------------+


// #### COMPATIBILITY OLD VERSION, TO BE REMOVED ####
- PROTMESSID_OPUS_SUPPORTED: Informs that OPUS codec is supported

    note: does not have any data -> n = 0


- PROTMESSID_RECORDER_STATE: notifies of changes in the server jam recorder state

    +--------------+
    | 1 byte state |
    +--------------+

    state is a value from the enum ERecorderState:
    - 0 undefined (not used by protocol messages)
    - tbc


CONNECTION LESS MESSAGES
------------------------

- PROTMESSID_CLM_PING_MS: Connection less ping message (for measuring the ping
                          time)

    +-----------------------------+
    | 4 bytes transmit time in ms |
    +-----------------------------+


- PROTMESSID_CLM_PING_MS_WITHNUMCLIENTS: Connection less ping message (for
                                         measuring the ping time) with the
                                         info about the current number of
                                         connected clients

    +-----------------------------+---------------------------------+
    | 4 bytes transmit time in ms | 1 byte number connected clients |
    +-----------------------------+---------------------------------+


- PROTMESSID_CLM_SERVER_FULL: Connection less server full message

    note: does not have any data -> n = 0


- PROTMESSID_CLM_REGISTER_SERVER: Register a server, providing server
                                  information

    +------------------------------+ ...
    | 2 bytes server internal port | ...
    +------------------------------+ ...
        ... -----------------+----------------------------------+ ...
        ...  2 bytes country | 1 byte maximum connected clients | ...
        ... -----------------+----------------------------------+ ...
        ... ---------------------+ ...
        ...  1 byte is permanent | ...
        ... ---------------------+ ...
        ... ------------------+----------------------------------+ ...
        ...  2 bytes number n | n bytes UTF-8 string server name | ...
        ... ------------------+----------------------------------+ ...
        ... ------------------+----------------------------------------------+ ...
        ...  2 bytes number n | n bytes UTF-8 string server internal address | ...
        ... ------------------+----------------------------------------------+ ...
        ... ------------------+---------------------------+
        ...  2 bytes number n | n bytes UTF-8 string city |
        ... ------------------+---------------------------+

    - "country" is according to "Common Locale Data Repository" which is used in
      the QLocale class
    - "maximum connected clients" is the maximum number of clients which can
      be connected to the server at the same time
    - "is permanent" is a flag which indicates if the server is permanent
      online or not. If this value is any value <> 0 indicates that the server
      is permanent online.
    - "server internal address" represents the IPv4 address as a dotted quad to
      be used by clients with the same external IP address as the server.
      NOTE: In the PROTMESSID_CLM_SERVER_LIST list, this field will be empty
      as only the initial IP address should be used by the client.  Where
      necessary, that value will contain the server internal address.


- PROTMESSID_CLM_REGISTER_SERVER_EX: Register a server, providing extended server
                                     information

    +--------------------------------+-------------------------------+
    | PROTMESSID_CLM_REGISTER_SERVER | PROTMESSID_CLM_VERSION_AND_OS |
    +--------------------------------+-------------------------------+


- PROTMESSID_CLM_UNREGISTER_SERVER: Unregister a server

    note: does not have any data -> n = 0


- PROTMESSID_CLM_SERVER_LIST: Server list message

    for each registered server append following data:

    +--------------------+--------------------------------+
    | 4 bytes IP address | PROTMESSID_CLM_REGISTER_SERVER |
    +--------------------+--------------------------------+

    - "PROTMESSID_CLM_REGISTER_SERVER" means that exactly the same message body
      of the PROTMESSID_CLM_REGISTER_SERVER message is used


- PROTMESSID_CLM_RED_SERVER_LIST: Reduced server list message (to have less UDP fragmentation)

    for each registered server append following data:

    +--------------------+------------------------------+ ...
    | 4 bytes IP address | 2 bytes server internal port | ...
    +--------------------+------------------------------+ ...
        ... -----------------+----------------------------------+
        ...  1 byte number n | n bytes UTF-8 string server name |
        ... -----------------+----------------------------------+


- PROTMESSID_CLM_REQ_SERVER_LIST: Request server list

    note: does not have any data -> n = 0


- PROTMESSID_CLM_SEND_EMPTY_MESSAGE: Send "empty message" message

    +--------------------+--------------+
    | 4 bytes IP address | 2 bytes port |
    +--------------------+--------------+


- PROTMESSID_CLM_DISCONNECTION: Disconnect message

    note: does not have any data -> n = 0


- PROTMESSID_CLM_VERSION_AND_OS: Version number and operating system

    +-------------------------+------------------+------------------------------+
    | 1 byte operating system | 2 bytes number n | n bytes UTF-8 string version |
    +-------------------------+------------------+------------------------------+


- PROTMESSID_CLM_REQ_VERSION_AND_OS: Request version number and operating system

    note: does not have any data -> n = 0


- PROTMESSID_CLM_CONN_CLIENTS_LIST: Information about connected clients

    for each connected client append the PROTMESSID_CONN_CLIENTS_LIST:

    +------------------------------+------------------------------+ ...
    | PROTMESSID_CONN_CLIENTS_LIST | PROTMESSID_CONN_CLIENTS_LIST | ...
    +------------------------------+------------------------------+ ...


- PROTMESSID_CLM_REQ_CONN_CLIENTS_LIST: Request the connected clients list

    note: does not have any data -> n = 0


- PROTMESSID_CLM_CHANNEL_LEVEL_LIST: The channel level list

    +----------------------------------+
    | ( ( n + 1 ) / 2 ) * 4 bit values |
    +----------------------------------+

    n is number of connected clients

    the values are the maximum channel levels for a client frame converted
    to the range of CLevelMeter in 4 bits, two entries per byte
    with the earlier channel in the lower half of the byte

    where an odd number of clients is connected, there will be four unused
    upper bits in the final byte, containing 0xF (which is out of range)

    the server may compute the message when any client has used
    PROTMESSID_CLM_REQ_CHANNEL_LEVEL_LIST to opt in

    the server should issue the message only to a client that has used
    PROTMESSID_CLM_REQ_CHANNEL_LEVEL_LIST to opt in


- PROTMESSID_CLM_REGISTER_SERVER_RESP: result of registration request

    +---------------+
    | 1 byte status |
    +---------------+

    - "status":
      Values of ESvrRegResult:
      0 - success
      1 - failed due to central server list being full
      2 - your server version is too old
      3 - registration requirements not fulfilled

    Note: the central server may send this message in response to a
          PROTMESSID_CLM_REGISTER_SERVER request.
          Where not received, the registering server may only retry up to
          five times for one registration request at 500ms intervals.
          Beyond this, it should "ping" every 15 minutes
          (standard re-registration timeout).
*/

package protocol

const (
	// Protocol Message ID
	PROTMESSID_ILLEGAL                  = 0
	PROTMESSID_ACKN                     = 1
	PROTMESSID_JITT_BUF_SIZE            = 10
	PROTMESSID_REQ_JITT_BUF_SIZE        = 11
	PROTMESSID_CHANNEL_GAIN             = 13
	PROTMESSID_REQ_CONN_CLIENTS_LIST    = 16
	PROTMESSID_CHAT_TEXT                = 18
	PROTMESSID_NETW_TRANSPORT_PROPS     = 20
	PROTMESSID_REQ_NETW_TRANSPORT_PROPS = 21
	PROTMESSID_REQ_CHANNEL_INFOS        = 23
	PROTMESSID_CONN_CLIENTS_LIST        = 24
	PROTMESSID_CHANNEL_INFOS            = 25
	PROTMESSID_OPUS_SUPPORTED           = 26
	PROTMESSID_LICENCE_REQUIRED         = 27
	PROTMESSID_VERSION_AND_OS           = 29
	PROTMESSID_CHANNEL_PAN              = 30
	PROTMESSID_MUTE_STATE_CHANGED       = 31
	PROTMESSID_CLIENT_ID                = 32
	PROTMESSID_RECORDER_STATE           = 33
	PROTMESSID_REQ_SPLIT_MESS_SUPPORT   = 34
	PROTMESSID_SPLIT_MESS_SUPPORTED     = 35

	// Protocol CLM
	PROTMESSID_CLM_PING_MS                = 1001
	PROTMESSID_CLM_PING_MS_WITHNUMCLIENTS = 1002
	PROTMESSID_CLM_SERVER_FULL            = 1003
	PROTMESSID_CLM_REGISTER_SERVER        = 1004
	PROTMESSID_CLM_UNREGISTER_SERVER      = 1005
	PROTMESSID_CLM_SERVER_LIST            = 1006
	PROTMESSID_CLM_REQ_SERVER_LIST        = 1007
	PROTMESSID_CLM_SEND_EMPTY_MESSAGE     = 1008
	PROTMESSID_CLM_EMPTY_MESSAGE          = 1009
	PROTMESSID_CLM_DISCONNECTION          = 1010
	PROTMESSID_CLM_VERSION_AND_OS         = 1011
	PROTMESSID_CLM_REQ_VERSION_AND_OS     = 1012
	PROTMESSID_CLM_CONN_CLIENTS_LIST      = 1013
	PROTMESSID_CLM_REQ_CONN_CLIENTS_LIST  = 1014
	PROTMESSID_CLM_CHANNEL_LEVEL_LIST     = 1015
	PROTMESSID_CLM_REGISTER_SERVER_RESP   = 1016
	PROTMESSID_CLM_REGISTER_SERVER_EX     = 1017
	PROTMESSID_CLM_RED_SERVER_LIST        = 1018

	// Splitted message ID
	PROTMESSID_SPECIAL_SPLIT_MESSAGE = 2001

	// Message information
	MESS_HEADER_LENGTH_BYTE    = 7
	MESS_LEN_WITHOUT_DATA_BYTE = MESS_HEADER_LENGTH_BYTE + 2
	MESS_SPLIT_PART_SIZE_BYTES = 550

	MAX_SIZE_BYTES_NETW_BUF = 20000
)

type Header struct {
	Tag    uint16
	ID     uint16
	Cnt    byte
	Length uint16
}

type Message struct {
	Hdr  Header
	CRC  uint16
	Data []byte
}
