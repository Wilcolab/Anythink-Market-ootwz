/**
 * Express router providing comment related routes.
 * @module routes/api/comments
 */

const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Get all comments.
 * @name GET/api/comments
 * @function
 * @memberof module:routes/api/comments
 * @inner
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {JSON} - A JSON object containing all comments
 */

/**
 * Delete a comment by ID.
 * @name DELETE/api/comments/:id
 * @function
 * @memberof module:routes/api/comments
 * @inner
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {JSON} - A JSON object indicating success
 */

module.exports = router;

router.get("/", async (req, res) => {
    try {
        const comments = await Comment.find();
        res.json({ comments });
    } catch (err) {
        console.log(err);
        res.status(500).send("Server Error");
    }
});

// add a delete endpoint for comments
router.delete("/:id", async (req, res) => {
    try {
        await Comment.findByIdAndRemove(req.params.id);
        res.json({ success: true });
    } catch (err) {
        console.log(err);
        res.status(500).send("Server Error");
    }
});
