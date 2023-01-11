import React, { useState } from 'react'
import { Col, Row } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'

const Home = () => {

    const [data, setdata] = useState(true)

    const selectData = (props) => {
        setdata(props);
    }

    return (
        <div>Hello</div>

    )
}

export default Home