import React from 'react';
import { nanoid } from 'nanoid';
import './css/styles.css';

class Main extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            roomId: "",
        }
    }

    createRoom = async (event) => {
        event.preventDefault()
        let roomId = nanoid(6)
        window.location.href = window.location.protocol + '//' + window.location.host + '/room/' + roomId;
    }

    joinRoom = async (event) => {
        event.preventDefault()
        if (this.state.roomId !==  '') {
            window.location.href = window.location.protocol + '//' + window.location.host + '/room/' + this.state.roomId
        }
    }
    render() {
        return (
            <div>
                <h1>Main Page</h1>
                <button onClick={this.createRoom}>Create Room</button>
                <form onSubmit={this.joinRoom}>
                    <input type={'text'} placeholder={'Room Code'} value={this.state.roomId}
                        onChange={(event) => this.setState({roomId: event.target.value})}/>
                    <button>Join Room</button>
                </form>
            </div>
        )
    }
}

export default Main;