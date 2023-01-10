import React, { useEffect, useState } from 'react'
import { Col, Row, Table } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'

import "./style.css"

const ChecklistPencarian = () => {

    return (
        <div>
            <Row className='w-100'>
                <Col xs={2}><Sidebars /></Col>
                <Col>
                    <Table striped hover >
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>First Name</th>
                                <th>Last Name</th>
                                <th>Username</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>Mark</td>
                                <td>Otto</td>
                                <td>@mdo</td>
                            </tr>
                            <tr>
                                <td>2</td>
                                <td>Jacob</td>
                                <td>Thornton</td>
                                <td>@fat</td>
                            </tr>
                            <tr>
                                <td>3</td>
                                <td colSpan={2}>Larry the Bird</td>
                                <td>@twitter</td>
                            </tr>
                        </tbody>
                    </Table>
                </Col>
            </Row>

        </div>
    )
}

export default ChecklistPencarian
