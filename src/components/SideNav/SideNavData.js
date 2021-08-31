import React from 'react';
import { FaHome, FaEnvelope, FaUserEdit } from 'react-icons/fa';
import { FiLogOut } from 'react-icons/fi';

export const SideNavData = [
    {
        title: "Home",
        icon: <FaHome />,
        link: '/',
    },
    {
        title: "Invites",
        icon: <FaEnvelope />,
        link: '/',
    },
    {
        title: "Edit Profile",
        icon: <FaUserEdit />,
        link: '/',
    },
    {
        title: "Logout",
        icon: <FiLogOut />,
        link: '/',
    }
]
