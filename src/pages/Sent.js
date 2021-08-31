import React from 'react';
import { MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBBtn } from 'mdb-react-ui-kit';
import './Invites.css';
import Toggle from '../components/InviteToggle.js';

function Sent() {
    return (
        <>
            <Toggle path='/sent' />

            <MDBCard className='w-75'>
                <MDBCardBody>
                    <MDBCardTitle>Card title</MDBCardTitle>
                    <MDBCardText className="text">With supporting text below as a natural lead-in to additional content.</MDBCardText>
                    <MDBBtn className='withdraw'>Withdraw</MDBBtn>
                </MDBCardBody>
            </MDBCard>

            <br />

            <MDBCard className='w-75'>
                <MDBCardBody>
                    <MDBCardTitle>Card title</MDBCardTitle>
                    <MDBCardText>With supporting text below as a natural lead-in to additional content.</MDBCardText>
                    <MDBBtn className='withdraw'>Withdraw</MDBBtn>
                </MDBCardBody>
            </MDBCard>

        </>
    );
}

export default Sent;