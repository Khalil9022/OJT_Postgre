import React, { useEffect, useState } from 'react'
import { Col, Dropdown, Row, Table } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'
import axios from "axios";
import { API_URL } from '../../utils/constant'

import "./style.css"

const ChecklistPencarian = () => {

    const [data, setdata] = useState(true)
    const [branch, setbranch] = useState()
    const [company, setcompany] = useState()
    const [opsibranch, setopsibranch] = useState()
    const [opsicompany, setopsicompany] = useState()

    const changeBranch = (data) => {
        setopsibranch(data)
    }

    const selectData = (props) => {
        setdata(props);
    }

    useEffect(() => {
        const fetchData = async () => {
            const result = await axios.get(API_URL + "branch");
            setbranch(result.data);

            const result2 = await axios.get(API_URL + "company");
            setcompany(result2.data);
        };

        fetchData()
    }, [])

    return (
        <div>
            <Row className='w-100'>
                <Col xs={data ? 2 : 1}><Sidebars selectData={selectData} /></Col>
                <Col className='table'>
                    <h1 className='text-center'>Data</h1>

                    <Dropdown>
                        <Dropdown.Toggle className="dr-down-toggle " variant="reds" id="dropdown-basic">
                            <span>{opsibranch?.data ? branch : "Pilih Kurir"}</span>

                        </Dropdown.Toggle>

                        <Dropdown.Menu>
                            {branch?.data.map((item, index) => (
                                <Dropdown.Item onClick={() => changeBranch(item.code)}>
                                    <div className="item-pics">
                                        {item.code}
                                    </div>
                                </Dropdown.Item>
                            ))}
                        </Dropdown.Menu>
                    </Dropdown>

                    <Table hover bordered={true} className='text-center' >
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
