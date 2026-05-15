import React, { useState, useEffect } from "react";

function Dashboard() {

  const [blogs, setBlogs] = useState([]);

  // Fetch blogs
  const fetchBlogs = async () => {
    const res = await fetch("http://localhost:30001/blogs");
    const data = await res.json();
    setBlogs(data);
  };

  useEffect(() => {
    fetchBlogs();
  }, []);

  return (
    <div className="container mt-4">

      <div className="row m-1 mt-4">
      <h2 className="mb-3">List of All Blogs</h2>

      {blogs.map((b) => (
        <div key={b.id} className="card mb-3 shadow-sm">
          <div className="card-body">
            <h4 className="card-title">{b.title}</h4>
            <p className="card-text">{b.content}</p>
          </div>
        </div>
      ))}
    </div>
    </div>
  );
}

export default Dashboard;