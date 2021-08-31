import React from 'react';
import { MDBCard, MDBCardBody, MDBCardTitle, MDBCardText, MDBBtn } from 'mdb-react-ui-kit';
import './Invites.css';
import Toggle from '../components/InviteToggle.js';

function Invites() {
    return (
        <>
            <Toggle path='/invites' />

            <MDBCard className='w-75'>
                <MDBCardBody>
                    <MDBCardTitle>Card title</MDBCardTitle>
                    <MDBCardText>With supporting text below as a natural lead-in to additional content.</MDBCardText>
                    <MDBBtn className='approve'>Approve</MDBBtn>
                    <MDBBtn className='decline'>Decline</MDBBtn>
                </MDBCardBody>
            </MDBCard>

            <br />

            <MDBCard className='w-75'>
                <MDBCardBody>
                    <MDBCardTitle>Card title</MDBCardTitle>
                    <MDBCardText>With supporting text below as a natural lead-in to additional content.</MDBCardText>
                    <MDBBtn className='approve'>Approve</MDBBtn>
                    <MDBBtn className='decline'>Decline</MDBBtn>
                </MDBCardBody>
            </MDBCard>

        </>
    );
}

export default Invites;