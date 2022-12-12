import React from "react";
import Button from "react-bootstrap/Button";
import InputGroup from "react-bootstrap/InputGroup";
import Form from "react-bootstrap/Form";
import { useNavigate } from "react-router-dom";
import { useMutation } from "react-query";
import Dummy from "../data/dummy";

function Add() {
  const navigate = useNavigate();

  const [form, setForm] = React.useState({
    id: new Date().getMilliseconds(),
    productID: Math.floor(Math.random() * 9999),
    productName: "",
    amount: "",
    customerName: "",
    status: 0,
    transactionDate: new Date().toString(),
    createBy: "",
    createOn: new Date().toString(),
  });

  const { productName, amount, customerName, status, createBy } = form;

  const handleChange = (e) => {
    setForm({
      ...form,
      [e?.target?.name]: e?.target?.value,
    });
  };

  const handleAdd = useMutation(async (e) => {
    try {
      e.preventDefault();
      Dummy.data.push(form);
      navigate("/");
    } catch (e) {
      console.log(e);
    }
  });

  return (
    <div>
      <div>
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Product Name"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={handleChange}
            name="productName"
            value={productName}
          />
        </InputGroup>
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Amount"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={handleChange}
            name="amount"
            value={amount}
          />
        </InputGroup>
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Customer Name"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={handleChange}
            name="customerName"
            value={customerName}
          />
        </InputGroup>
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Status"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={handleChange}
            name="status"
            value={status}
          />
        </InputGroup>
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Created By"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={handleChange}
            name="createBy"
            value={createBy}
          />
        </InputGroup>
        <Button variant="primary" onClick={(e) => handleAdd.mutate(e)}>
          Add Data
        </Button>
      </div>
    </div>
  );
}

export default Add;
