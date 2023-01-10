import React, { useEffect, useState } from 'react'
import { Col, Row, Table } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'

import "./style.css"

const ChecklistPencarian = () => {

    const [data, setdata] = useState(true)

    const selectData = (props) => {
        setdata(props);
    }

    return (
        <div>
            <Row className='w-100'>
                <Col xs={data ? 2 : 1}><Sidebars selectData={selectData} /></Col>
                <Col className='table'>
                    <h1 className='text-center'>Data</h1>
                    <Table striped hover  >
                        <thead>
                            <tr>
                                <th>PPK</th>
                                <th>Name</th>
                                <th>Channeling Company</th>
                                <th>Drawdown Date</th>
                                <th>Loan Amount</th>
                                <th>Loan Period</th>
                                <th>Interest Eff</th>
                                <th>Check</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>Mark</td>
                                <td>Otto</td>
                                <td>@mdo</td>
                                <td>1</td>
                                <td>Mark</td>
                                <td>Otto</td>
                                <td>@mdo</td>
                            </tr>
                        </tbody>
                    </Table>
                </Col>
            </Row>

        </div>
    )
}

export default ChecklistPencarian
