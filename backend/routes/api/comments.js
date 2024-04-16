const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Express router for handling comments API.
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
 * @returns {Object} - JSON response containing the comments.
 * @throws {Object} - JSON response containing the error message.
 */
router.get("/", (req, res) => {
    Comment.find()
        .then(comments => {
            res.json({ comments });
        })
        .catch(err => {
            res.status(500).json({ error: err.message });
        });
});

/**
 * DELETE /api/comments/:id
 * Deletes a comment by ID.
 * @name DELETE /api/comments/:id
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object.
 * @param {Object} res - Express response object.
 * @returns {Object} - JSON response indicating the success of the deletion.
 * @throws {Object} - JSON response containing the error message.
 */
router.delete("/:id", async (req, res) => {
    try {
        await Comment.findByIdAndRemove(req.params.id);
        res.json({ success: true });
    } catch (err) {
        res.status(500).json({ error: err.message });
    }
});
