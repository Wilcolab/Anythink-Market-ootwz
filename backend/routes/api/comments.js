const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Express router for handling comments API requests.
 * @module routes/api/comments
 */

module.exports = router;

/**
 * Route for getting all comments.
 * @name GET /api/comments
 * @function
 * @async
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {Object} - JSON response containing all comments
 */
router.get("/", async (req, res) => {
    const comments = await Comment.find();
    res.json(comments);
});

/**
 * Route for deleting a comment by ID.
 * @name DELETE /api/comments/:id
 * @function
 * @async
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {Object} - JSON response indicating success
 */
router.delete("/:id", async (req, res) => {
    const comment = await Comment.findById(req.params.id);
    await comment.remove();
    res.json({ success: true });
});
