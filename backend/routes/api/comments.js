/**
 * Express router for handling comment-related API routes.
 * @module routes/api/comments
 */

const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

/**
 * Route for retrieving all comments.
 * @name GET /api/comments
 * @function
 * @async
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @param {Function} next - Express next middleware function
 * @returns {Object} - JSON response containing the retrieved comments
 */
router.get("/", async (req, res, next) => {
    try {
        const comments = await Comment.find();
        return res.json({ comments: comments });
    } catch (error) {
        next(error);
    }
});

/**
 * Route for deleting a comment by its ID.
 * @name DELETE /api/comments/:id
 * @function
 * @async
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @param {Function} next - Express next middleware function
 * @returns {Object} - HTTP response with status 200 if successful
 */
router.delete("/:id", async (req, res, next) => {
    try {
        await Comment.findByIdAndDelete(req.params.id);
        return res.sendStatus(200);
    } catch (error) {
        next(error);
    }
});
