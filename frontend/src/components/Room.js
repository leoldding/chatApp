import React from 'react';
import './css/styles.css';
import './css/room.css';

class Room extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            inputText: "",
        };
    }

    sendMessage = async (event) => {
        event.preventDefault();
        console.log(this.state.inputText);
        this.setState({inputText: ""});
    }
    render() {
        return (
            <div>
                <h1>Chat Room</h1>
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