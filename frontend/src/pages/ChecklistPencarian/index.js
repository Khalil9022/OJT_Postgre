import React, { useEffect, useState } from 'react'
import { Button, Col, Container, Dropdown, Form, Row, Table } from 'react-bootstrap'
import Sidebars from '../../components/Sidebar'
import axios from "axios";
import { API_URL } from '../../utils/constant'

import "./style.css"

const ChecklistPencarian = () => {

    const [branch, setbranch] = useState()
    const [company, setcompany] = useState()
    const [opsibranch, setopsibranch] = useState()
    const [opsicompany, setopsicompany] = useState()
    const [datestart, setdatestart] = useState()
    const [dateend, setdateend] = useState()
    const [allData, setallData] = useState()

    const changeBranch = (data) => {
        setopsibranch(data)
    }

    const changeCompany = (data) => {
        setopsicompany(data)
    }

    const changedatestart = (data) => {
        setdatestart(data)
    }

    const changedateend = (data) => {
        setdateend(data)
    }

    useEffect(() => {
        const fetchData = async () => {
            const result = await axios.get(API_URL + "branch");
            setbranch(result.data);

            const result2 = await axios.get(API_URL + "company");
            setcompany(result2.data);

            const result3 = await axios.get(API_URL + "allcustomer");
            setallData(result3.data);
        };

        fetchData()
    }, [])

    return (
        <div >
            <Container fluid>
                <Row>

                    <Col className='table'>
                        <h1 className='text-center mb-3 mt-5'>Data</h1>

                        <Row className='mb-4 '>
                            <Col className='d-flex align-items-center ms-5' xs={3}>
                                <span>Branch</span>
                                <Form.Select size="sm">
                                    <option>{opsibranch ? opsibranch : "Select Branch"}</option>
                                    {branch?.data.map((item, index) => (
                                        <option onClick={() => changeBranch(item.code)}>
                                            {item.code}
                                        </option>
                                    ))}
                                </Form.Select>
                            </Col>

                            <Col className='d-flex align-items-center' xs={3}>
                                <span>Company</span>
                                <Form.Select size="sm">
                                    <option>{opsicompany ? opsicompany : "Select Company"}</option>
                                    {company?.data.map((item, index) => (
                                        <option onClick={() => changeCompany(item.company_short_name)}>
                                            {item.company_short_name}
                                        </option>
                                    ))}
                                </Form.Select>
                            </Col>

                            <Col className='d-flex align-items-center' xs={2}>
                                <span>Start</span>
                                <Form.Control type="date" onChange={(event) => changedatestart(event.target.value)} />
                            </Col>
                            <Col className='d-flex align-items-center' xs={2}>

                                <span>End</span>
                                <Form.Control type="date" onChange={(event) => changedateend(event.target.value)} />
                            </Col>

                            <Col className='d-flex align-items-center'>
                                <Button className='btn btn-outline-primary btn-sm'>
                                    Submit
                                </Button>
                            </Col>

                        </Row>

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
                                {allData?.data.map((item, index) => (
                                    <>
                                        <tr key={index}>
                                            <td>{item.PPK}</td>
                                            <td>{item.Name}</td>
                                            <td>{item.ChannelingCompany}</td>
                                            <td>{item.DrawdownDate.substring(0, 10)}</td>
                                            <td>{item.LoanAmount}</td>
                                            <td>{item.LoanPeriod}</td>
                                            <td>{item.InterestEffective}</td>
                                            <td><Form.Check name="check" /></td>
                                        </tr>
                                    </>
                                ))}
                            </tbody>
                        </Table>
                    </Col>
                </Row>
            </Container>
        </div>
    )
}

export default ChecklistPencarian
