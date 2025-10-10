/**
 * GET /api/comments/
 * Retrieves all comments from the database.
 * 
 * @route GET /api/comments/
 * @returns {Array<Object>} 200 - An array of comment objects
 * @returns {Error} 500 - Failed to fetch comments
 */

/**
 * DELETE /api/comments/:id
 * Deletes a comment by its ID.
 * 
 * @route DELETE /api/comments/:id
 * @param {string} id.path.required - The ID of the comment to delete
 * @returns {Object} 200 - Success message
 * @returns {Error} 404 - Comment not found
 * @returns {Error} 500 - Failed to delete comment
 */
 * Express router for handling comment-related API routes.
 * @type {import('express').Router}
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

// add a delete endpoint
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
