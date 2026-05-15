import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Home from "./Home";
import Login from "./Login";
import CreateBlog from "./Blogs";
import Dashboard from "./Dashboard";
import ImageUpload from "./ImageUpload";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/blogs" element={<CreateBlog />} />
        <Route path="/dashboard" element={<Dashboard />} />
	<Route path="/fileupload" element={<ImageUpload />} />
      </Routes>
    </Router>
  );
}

export default App;