import React, { useState } from 'react'
import * as FaIcons from "react-icons/fa";
import * as AiIcons from "react-icons/ai";
import { Link } from 'react-router-dom';
import { SidebarData } from './sidebardata';
import './style.css'
import { IconContext } from 'react-icons';



const Sidebars = ({ selectData }) => {
    const [sidebar, setsidebar] = useState(true)

    const showSidebar = () => {
        setsidebar(!sidebar)
        selectData(!sidebar)
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
                    <ul className='nav-menu-items' onClick={showSidebar}>
                        <li className='navbar-toggle'>
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
                        <li className=' nav-text position-absolute bottom-0'>
                            <Link to='/'>
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