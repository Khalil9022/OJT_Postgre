import React from 'react'
import * as BsIcons from 'react-icons/bs'
import * as AiIcons from 'react-icons/ai'
import * as IoIcons from 'react-icons/io'

export const SidebarData = [
    {
        title: 'Home',
        path: '/',
        icon: <AiIcons.AiFillHome />,
        cName: 'nav-text'
    },
    {
        title: 'Checklist Pencairan',
        path: '/checklistpencairan',
        icon: <BsIcons.BsCardChecklist />,
        cName: 'nav-text'
    },
    {
        title: 'Drawdown Report',
        path: '/report',
        icon: <IoIcons.IoIosPaper />,
        cName: 'nav-text'
    }
]
