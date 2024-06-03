const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Express router for handling comment-related API endpoints.
 * @module routes/api/comments
 */

module.exports = router;

/**
 * GET /api/comments
 * Retrieves all comments.
 * @name GET /api/comments
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object.
 * @param {Object} res - Express response object.
 * @returns {Object} - JSON response containing all comments.
 */
router.get("/", async (req, res) => {
  const comments = await Comment.find();
  res.json(comments);
});

/**
 * DELETE /api/comments/:id
 * Deletes a comment by its ID.
 * @name DELETE /api/comments/:id
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object.
 * @param {Object} res - Express response object.
 * @returns {Object} - Empty response with status 204 if successful, or error response with status 500 if failed.
 */
router.delete("/:id", async (req, res) => {
    try {
        const { id } = req.params;
        await Comment.findByIdAndDelete(id);
        res.status(204).send();
    } catch (error) {
        res.status(500).json({ error: "Failed to delete comment" });
    }
});
