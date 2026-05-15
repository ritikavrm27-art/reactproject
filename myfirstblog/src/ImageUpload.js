import React, { useState } from "react";

function ImageUpload() {
  const [file, setFile] = useState(null);
  const [preview, setPreview] = useState(null);
  const [uploadedUrl, setUploadedUrl] = useState("");

  // handle file select
  const handleChange = (e) => {
    const selectedFile = e.target.files[0];
    setFile(selectedFile);

    // create preview
    const previewUrl = URL.createObjectURL(selectedFile);
    setPreview(previewUrl);
  };

  // upload to backend
  const handleUpload = async () => {
    if (!file) {
      alert("Select a file first");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const res = await fetch("http://localhost:30001/upload", {
        method: "POST",
        body: formData
      });

      const data = await res.json();

      // backend should return file URL
      setUploadedUrl(data.url);

    } catch (err) {
      console.error(err);
      alert("Upload failed");
    }
  };

  return (
    <div className="container mt-4">
      <h3>Upload Image</h3>

      <input type="file" accept="image/*" onChange={handleChange} />

      {/* Preview BEFORE upload */}
      {preview && (
        <div className="mt-3">
          <p>Preview:</p>
          <img src={preview} alt="preview" width="200" />
        </div>
      )}

      <button className="btn btn-primary mt-3" onClick={handleUpload}>
        Upload
      </button>

      {/* Show uploaded image */}
      {uploadedUrl && (
        <div className="mt-3">
          <p>Saved Image:</p>
          <img src={uploadedUrl} alt="uploaded" width="200" />
        </div>
      )}
    </div>
  );
}

export default ImageUpload;