import React, { useState } from 'react'
import * as FaIcons from "react-icons/fa";
import * as AiIcons from "react-icons/ai";
import { Link, useNavigate } from 'react-router-dom';
import { SidebarData } from './sidebardata';
import './style.css'
import { IconContext } from 'react-icons';
import swal from 'sweetalert';



const Sidebars = ({ selectData }) => {
    const [sidebar, setsidebar] = useState(true)

    const navigate = useNavigate()

    const showSidebar = () => {
        setsidebar(!sidebar)
        selectData(!sidebar)
    }

    const handleLogout = () => {
        localStorage.removeItem("isLoggedin")
        localStorage.removeItem("token")
        swal("Sukses", "Berhasil Logout", "success").then(
            navigate("/"),
            window.location.reload()
        )

    }

    return (
        <>
            {/* <div className='navbar'>
                <Link to='#' className='menu-bars'>
                    <FaIcons.FaBars onClick={showSidebar} />
                </Link>
            </div> */}
            <IconContext.Provider value={{ color: '#fff' }}>
                <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
                    <ul className='nav-menu-items' >
                        <li className='navbar-toggle' onClick={showSidebar}>
                            <Link to='#' className='menu-bars'>
                                <FaIcons.FaBars />
                            </Link>
                        </li>
                        {SidebarData.map((item, index) => {
                            return (
                                <li key={index} className={item.cName}>
                                    <Link to={item.path}>
                                        {item.icon}
                                        <span >{item.title}</span>
                                    </Link>
                                </li>
                            );
                        })}
                    </ul>
                    <div>
                        <li className='nav-text position-absolute bottom-0' >
                            <Link onClick={handleLogout}>
                                <FaIcons.FaDoorClosed />
                                <span >Keluar</span>
                            </Link>
                        </li>
                    </div>
                </nav>
            </IconContext.Provider>
        </>
    )
}

export default Sidebars 