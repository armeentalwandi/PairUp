import React, { useState } from 'react'
import { FaBars } from 'react-icons/fa';
import { AiOutlineClose } from 'react-icons/ai';
import { SideNavData } from './SideNavData';
import { Link } from 'react-router-dom';
import './SideNav.css';


function Navbar() {
    const [isOpened, setIsOpened] = useState(false);

    const showSideNav = () => { setIsOpened(!isOpened) };

    return (
        <>
            <div className='sidenav'>
                <Link to='#' className='menu-bars'>
                    <FaBars onClick={showSideNav} />
                </Link>
            </div>
            <nav className={isOpened ? 'opened-nav' : 'closed-nav'}>
                <ul className='menu-items' onClick={showSideNav}>
                    <li className='close-nav'>
                        <Link to='#' className='menu-bars'>
                            <AiOutlineClose />
                        </Link>
                    </li>
                    {SideNavData.map((item, index) => {
                        return (
                            <li key={index} className='item-text'>
                                <Link to={item.link}>
                                    {item.icon}
                                    <span>{item.title}</span>
                                </Link>
                            </li>
                        );
                    })}
                </ul>
            </nav>
        </>
    );
}

export default Navbar;