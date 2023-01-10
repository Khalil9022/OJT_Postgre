import React, { useState } from 'react'
import { Col, Row } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'

const Home = () => {

    const [data, setdata] = useState(true)

    const selectData = (props) => {
        setdata(props);
    }

    return (
        <Row>
            <Col xs={data ? 2 : 1}><Sidebars selectData={selectData} /></Col>
            <Col>Hello</Col>
        </Row>

    )
}

export default Home