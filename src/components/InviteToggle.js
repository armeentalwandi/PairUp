import React from 'react';
import '../pages/Invites.css';
import { Link } from "react-router-dom";

const Toggle = (props) => {
    return (<div className='toggle'>
        <Link to='/invites' className={props.path === '/invites' ? 'toggle-item-here' : 'toggle-item'} >
            Recieved
        </Link>

        <span> <Link to='/sent' className={props.path === '/sent' ? 'toggle-item-here' : 'toggle-item'}>
            Sent
        </Link> </span>
        <hr class="solid"></hr>
    </div>);
}

export default Toggle;