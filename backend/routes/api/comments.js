/**
 * @module routes/api/comments
 * @description Express router for handling comment-related API endpoints.
 */

 /**
    * GET /
    * Retrieves all comments, sorted by creation date (most recent first).
    * @route GET /api/comments
    * @group Comments - Operations about comments
    * @returns {Array.<Comment>} 200 - An array of comment objects
    * @returns {object} 500 - Internal server error
    */

 /**
    * DELETE /:id
    * Deletes a comment by its unique identifier.
    * @route DELETE /api/comments/{id}
    * @group Comments - Operations about comments
    * @param {string} id.path.required - The ID of the comment to delete
    * @returns {object} 200 - Success message
    * @returns {object} 404 - Comment not found
    * @returns {object} 500 - Internal server error
    */
const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

router.get("/", async (req, res) => {
  try {
    const comments = await Comment.find().sort({ createdAt: -1 });
    res.status(200).json(comments);
  } catch (error) {
    console.error("Error fetching comments:", error);
    res.status(500).json({ message: "Internal server error" });
  }
});

// add a delete endpoint to delete a comment by id
router.delete("/:id", async (req, res) => {
  const { id } = req.params;
  try {
    const comment = await Comment.findByIdAndDelete(id);
    if (!comment) {
      return res.status(404).json({ message: "Comment not found" });
    }
    res.status(200).json({ message: "Comment deleted successfully" });
  } catch (error) {
    console.error("Error deleting comment:", error);
    res.status(500).json({ message: "Internal server error" });
  }
}