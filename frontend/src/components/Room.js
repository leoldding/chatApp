import React from 'react';
import './css/styles.css';
import './css/room.css';

class Room extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            inputText: "",
            roomId: "",
        };
    }

    ws = new WebSocket('ws://' + window.location.host + '/ws' + window.location.pathname)

    async componentDidMount() {
        let log = document.getElementById("chatLog")
        let splitURL = window.location.href.split("/")
        this.setState({roomId: splitURL[splitURL.length - 1]})
        this.ws.onmessage = event => {
            let msg = document.createElement("div")
            msg.innerHTML = event.data
            log.appendChild(msg)
            log.scrollTop = log.scrollHeight - log.clientHeight
        }

        this.ws.onerror = event => {
            let msg = document.createElement("div")
            msg.innerHTML = "Connection has been closed."
            log.appendChild(msg)
            log.scrollTop = log.scrollHeight - log.clientHeight
        }

        this.ws.onclose = event => {
            let msg = document.createElement("div")
            msg.innerHTML = "Connection has been closed."
            log.appendChild(msg)
            log.scrollTop = log.scrollHeight - log.clientHeight
        }
    }


    sendMessage = async (event) => {
        event.preventDefault()
        if (this.state.inputText !== "") {
            this.ws.send(this.state.inputText)

            this.setState({inputText: ""})
        }
    }

    render() {
        let roomId = this.state.roomId
        return (
            <div>
                <h1>Chat Room</h1>
                <p>Room Code: {roomId}</p>
                <div id={"chatLog"}></div>
                <form onSubmit={this.sendMessage}>
                    <input type={"text"} placeholder={"Message"} value={this.state.inputText}
                           onChange={(event) => this.setState({inputText: event.target.value})}/>
                    <button>Send</button>
                </form>
            </div>
        )
    }

}

export default Room;