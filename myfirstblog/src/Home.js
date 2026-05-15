import React from "react";
import { Link } from "react-router-dom";

function Home() {
  return (
    <div className="container text-center mt-5">
      <h1>Welcome to home page</h1>

      <div className="mt-4">
        <Link to="/login" className="btn btn-primary m-2">
          Login
        </Link>

        <Link to="/dashboard" className="btn btn-success m-2">
          View Dashbaord
        </Link>


	<Link to="/fileupload" className="btn btn-warning m-2">
          Upload File
        </Link>
      </div>
    </div>
  );
}

export default Home;