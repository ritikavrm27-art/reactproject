import React, { useState, useEffect } from "react";

function CreateBlog() {
  const [formData, setFormData] = useState({
    title: "",
    content: "",
    author: ""
  });

  const [blogs, setBlogs] = useState([]);

  // Fetch blogs
  const fetchBlogs = async () => {
    const res = await fetch("http://localhost:8001/blogs");
    const data = await res.json();
    setBlogs(data);
  };

  useEffect(() => {
    fetchBlogs();
  }, []);

  // Handle input
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  // Submit form
  const handleSubmit = async (e) => {
    e.preventDefault();

    await fetch("http://localhost:8001/blog/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(formData)
    });

    setFormData({ title: "", content: "", author: "" });

    // Refresh list
    fetchBlogs();
  };

  return (
    <div className="container mt-4">
  <div className="card shadow p-4">
    <h2 className="mb-4">Create Blog</h2>

    <form onSubmit={handleSubmit}>

      {/* Title */}
      <div className="mb-3">
        <label className="form-label">Title</label>
        <input
          className="form-control"
          name="title"
          placeholder="Enter title"
          value={formData.title}
          onChange={handleChange}
        />
      </div>

      {/* Content */}
      <div className="mb-3">
        <label className="form-label">Content</label>
        <textarea
          className="form-control"
          rows="4"
          name="content"
          placeholder="Write your blog..."
          value={formData.content}
          onChange={handleChange}
        />
      </div>

      {/* Author */}
      <div className="mb-3">
        <label className="form-label">Author</label>
        <input
          className="form-control"
          name="author"
          placeholder="Author name"
          value={formData.author}
          onChange={handleChange}
        />
      </div>

      {/* Button */}
      <button type="submit" className="btn btn-primary w-100">
        Create Blog
      </button>

    </form>
  </div>

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

export default CreateBlog;