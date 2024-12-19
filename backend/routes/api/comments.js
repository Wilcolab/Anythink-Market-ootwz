/**
 * Express router providing comment related routes.
 * @module routes/api/comments
 */

const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Get all comments.
 * @name GET /
 * @function
 * @memberof module:routes/api/comments
 * @inner
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {JSON} - A JSON array of comments
 */

/**
 * Delete a comment by ID.
 * @name DELETE /:commentId
 * @function
 * @memberof module:routes/api/comments
 * @inner
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {StatusCode} 204 - No Content
 */

module.exports = router;

router.get("/", async (req, res) => {
    try {
        const comments = await Comment.find();
        res.json(comments);
    } catch (err) {
        console.error(err);
        res.status(500).send("Internal Server Error");
    }
});

router.delete("/:commentId", async (req, res) => {
    try {
        await Comment.findByIdAndRemove(req.params.commentId);
        res.sendStatus(204);
    } catch (err) {
        console.error(err);
        res.status(500).send("Internal Server Error");
    }
});