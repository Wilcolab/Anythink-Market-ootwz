/**
 * Express router for handling comment-related API endpoints.
 * 
 * @module routes/api/comments
 */

 /**
  * GET /
  * Retrieves all comments from the database.
  * 
  * @route GET /api/comments
  * @returns {Object[]} 200 - An array of comment objects
  * @returns {Object} 500 - Error message if fetching fails
  */

 /**
  * DELETE /:id
  * Deletes a comment by its ID.
  * 
  * @route DELETE /api/comments/:id
  * @param {string} id - The ID of the comment to delete
  * @returns {Object} 200 - Success message if deleted
  * @returns {Object} 404 - Error message if comment not found
  * @returns {Object} 500 - Error message if deletion fails
  */
const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

router.get("/", async (req, res) => {
  try {
    const comments = await Comment.find();
    res.json(comments);
  } catch (err) {
    res.status(500).json({ error: "Failed to fetch comments" });
  }
});

// add another endpoint for deleting a comment
router.delete("/:id", async (req, res) => {
  try {
    const comment = await Comment.findByIdAndDelete(req.params.id);
    if (!comment) {
      return res.status(404).json({ error: "Comment not found" });
    }
    res.json({ message: "Comment deleted successfully" });
  } catch (err) {
    res.status(500).json({ error: "Failed to delete comment" });
  }
});