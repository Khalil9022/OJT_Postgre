import React, { useEffect, useState } from 'react'
import { Button, Col, Container, Form, Row, Table } from 'react-bootstrap'
import axios from "axios";
import { API_URL } from '../../utils/constant'

import "./style.css"
import swal from 'sweetalert';

const ChecklistPencarian = () => {

    const [branch, setbranch] = useState()
    const [company, setcompany] = useState()
    const [opsibranch, setopsibranch] = useState("000")
    const [opsicompany, setopsicompany] = useState("000")
    const [datestart, setdatestart] = useState("")
    const [dateend, setdateend] = useState("")
    const [allData, setallData] = useState()
    const [checked, setchecked] = useState([])
    const [currentdate, setcurrentdate] = useState()

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

    const handleSubmit = () => {
        if (datestart !== "" && dateend !== "") {

            const apifilter = async () => {
                await axios.post(API_URL + "spesifikcustomer", {
                    branch: opsibranch,
                    company: opsicompany,
                    start: datestart,
                    end: dateend
                }).then((result2) => {
                    setallData(result2.data)
                }).catch((error) => {
                    console.log("Error, tidak dapat parsing data ke API", error);
                })
            }

            apifilter()

            swal("Sukses", "Sukses Melakukan Filter", "success");
        } else {
            swal("Gagal", "Silahkan isi input datenya terlebih dahulu", "warning");
        }

    }

    const handleChange = (event, ppk) => {
        // console.log(event.target.checked);
        // console.log(ppk);
        if (event.target.checked == true) {
            setchecked((i) => [...i, ppk])
        } else {
            setchecked((i) => i.filter((j) => j != ppk))
        }

    }

    const handleApprove = () => {
        console.log(checked);
        const approvalchange = async () => {
            await axios.post(API_URL + "updateapproval", {
                ppk: checked
            }).then(async () => {
                const alldata = await axios.get(API_URL + "allcustomer");
                setallData(alldata.data);
                swal("Sukses", "Sukses Melakukan Approve", "success");
            }).catch((error) => {
                console.log("Error, tidak dapat parsing data ke API", error);
                swal("Gagal", "Gagal melakukan approve", "warning");
            })
        }

        approvalchange()
        // document.querySelectorAll('input[type=checkbox]').forEach(el => el.checked = false);
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

        setcurrentdate(new Date().toLocaleString('en-us', { year: 'numeric', month: '2-digit', day: '2-digit' }).replace(/(\d+)\/(\d+)\/(\d+)/, '$3-$1-$2'))

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
                                <Form.Select size="sm" onChange={(e) => changeBranch(e.target.value)}>
                                    <option>{opsibranch ? "000 - AllBranch" : "Select Branch"}</option>
                                    {branch?.data.map((item, index) => (
                                        <option value={item.code}>
                                            00{item.code}
                                        </option>
                                    ))}
                                </Form.Select>
                            </Col>

                            <Col className='d-flex align-items-center' xs={3}>
                                <span>Company</span>
                                <Form.Select size="sm" onChange={(e) => changeCompany(e.target.value)}>
                                    <option>{opsicompany ? "AllCompany" : "Select Company"}</option>
                                    {company?.data.map((item, index) => (
                                        <option value={item.company_short_name}>
                                            {item.company_short_name}
                                        </option>
                                    ))}
                                </Form.Select>
                            </Col>

                            <Col className='d-flex align-items-center' xs={2}>
                                <span>Start</span>
                                <Form.Control type="date" defaultValue={currentdate} onChange={(event) => changedatestart(event.target.value)} />
                            </Col>
                            <Col className='d-flex align-items-center' xs={2}>

                                <span>End</span>
                                <Form.Control type="date" defaultValue={currentdate} onChange={(event) => changedateend(event.target.value)} />
                            </Col>

                            <Col className='d-flex align-items-center'>
                                <Button className='btn btn-outline-primary btn-sm' onClick={handleSubmit}>
                                    Submit
                                </Button>
                            </Col>

                        </Row>

                        <Table hover striped bordered={true} size="md">
                            <thead style={{ color: "white", background: "black" }}>
                                <tr >

                                    <th>PPK</th>
                                    <th>Name</th>
                                    <th>Channeling Company</th>
                                    <th>Drawdown Date</th>
                                    <th>Loan Amount</th>
                                    <th>Loan Period</th>
                                    <th>Interest Eff</th>
                                    <th></th>
                                </tr>
                            </thead>
                            <tbody>

                                {allData?.data ? allData.data.map((item, index) => (
                                    <>
                                        <tr key={index}>
                                            <td>{item.PPK}</td>
                                            <td>{item.Name}</td>
                                            <td>{item.ChannelingCompany}</td>
                                            <td>{item.DrawdownDate.substring(0, 10)}</td>
                                            <td>{item.LoanAmount}</td>
                                            <td>{item.LoanPeriod}</td>
                                            <td>{item.InterestEffective}%</td>
                                            <td><Form.Check name="check" checked={checked.includes(item.PPK) ? true : false} onChange={(e) => handleChange(e, item.PPK)} /></td>
                                        </tr>
                                    </>
                                ))

                                    : <td className='text-center border' colSpan={8}><b>Tidak ada Data</b></td>
                                }
                            </tbody>
                        </Table>
                        <Button className='btn btn-success btn-sm' onClick={handleApprove}>
                            Approve
                        </Button>
                    </Col>
                </Row>
            </Container>
        </div>
    )
}

export default ChecklistPencarian
