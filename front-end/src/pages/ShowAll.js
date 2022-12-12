import Table from "react-bootstrap/Table";
import React from "react";
import Button from "react-bootstrap/Button";
import InputGroup from "react-bootstrap/InputGroup";
import Form from "react-bootstrap/Form";
import { useNavigate } from "react-router-dom";
import Dummy from "../data/dummy";

function ShowAll() {
  const [query, setQuery] = React.useState("");
  const navigate = useNavigate();

  return (
    <div>
      <div className="d-flex justify-content-between">
        <InputGroup style={{ width: "30%" }} className="mb-3">
          <Form.Control
            placeholder="Search Product Name/Customer Name"
            aria-label="Search"
            aria-describedby="basic-addon1"
            style={{
              backgroundColor: "#E7E7E7",
              borderStyle: "none",
              borderRadius: "6px",
            }}
            onChange={(e) => setQuery(e?.target?.value)}
          />
        </InputGroup>
        <Button
          className="mb-3"
          variant="primary"
          onClick={() => {
            navigate(`/add`);
          }}
        >
          Add Data
        </Button>
      </div>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>ID</th>
            <th>Product Name</th>
            <th>Amount</th>
            <th>Customer Name</th>
            <th>Transaction Date</th>
            <th>Status</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {Dummy?.data
            ?.filter((o) => {
              return query.toLocaleLowerCase() === ""
                ? o
                : o?.productName?.toLocaleLowerCase().includes(query) ||
                    o?.customerName?.toLocaleLowerCase().includes(query);
            })
            .map((p, index) => (
              <tr key={p?.id}>
                <td>{p?.id}</td>
                <td
                  onClick={() => {
                    navigate(`/detail/${index}`);
                  }}
                >
                  {p?.productName}
                </td>
                <td>{p?.amount}</td>
                <td>{p?.customerName}</td>
                <td>{p?.transactionDate}</td>
                <td>
                  {p?.status === 0
                    ? "SUCCESS"
                    : p?.status === 1
                    ? "FAILED"
                    : ""}
                </td>
                <td className="d-flex justify-content-evenly">
                  <Button
                    variant="primary"
                    onClick={() => {
                      navigate(`/edit/${index}`);
                    }}
                  >
                    Edit
                  </Button>
                  <Button
                    variant="primary"
                    onClick={() => {
                      Dummy?.data?.splice(index, 1);
                    }}
                  >
                    Delete
                  </Button>
                </td>
              </tr>
            ))}
        </tbody>
      </Table>
    </div>
  );
}

export default ShowAll;
