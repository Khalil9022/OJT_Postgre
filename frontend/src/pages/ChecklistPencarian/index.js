import React from 'react'
import { Menu, MenuItem, Sidebar, useProSidebar } from 'react-pro-sidebar'
import Sidebars from '../../components/Sidebar'

import "./style.css"

const ChecklistPencarian = () => {
    const { collapseSidebar } = useProSidebar();

    return (
        <div style={{ display: 'flex', height: '100%' }}>
            <Sidebar>
                <Menu>
                    <MenuItem> Documentation</MenuItem>
                    <MenuItem> Calendar</MenuItem>
                    <MenuItem> E-commerce</MenuItem>
                </Menu>
            </Sidebar>
            <main>
                <button onClick={() => collapseSidebar()}>Collapse</button>
            </main>
        </div>
    )
}

export default ChecklistPencarian
