import Table from "react-bootstrap/Table";
import React from "react";
import Button from "react-bootstrap/Button";
import { useNavigate } from "react-router-dom";
import Dummy from "../data/dummy";
import { useParams } from "react-router-dom";

function Detail() {
  const navigate = useNavigate();
  const [query, setQuery] = React.useState("");
  let { id } = useParams();

  let dummy2 = Dummy.data[id];

  return (
    <div>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>ID</th>
            <th>Product ID</th>
            <th>Product Name</th>
            <th>Amount</th>
            <th>Customer Name</th>
            <th>Transaction Date</th>
            <th>Status</th>
            <th>Created By</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{dummy2?.id}</td>
            <td>{dummy2?.productID}</td>
            <td>{dummy2?.productName}</td>
            <td>{dummy2?.amount}</td>
            <td>{dummy2?.customerName}</td>
            <td>{dummy2?.transactionDate}</td>
            <td>
              {dummy2?.status === 0
                ? "SUCCESS"
                : dummy2?.status === 1
                ? "FAILED"
                : ""}
            </td>
            <td>{dummy2?.createBy}</td>
            <td className="d-flex justify-content-evenly">
              <Button
                variant="primary"
                onClick={() => {
                  navigate(`/edit/${id}`);
                }}
              >
                Edit
              </Button>
              <Button
                variant="primary"
                onClick={() => {
                  Dummy.data.splice(id, 1);
                }}
              >
                Delete
              </Button>
            </td>
          </tr>
        </tbody>
      </Table>
    </div>
  );
}

export default Detail;
