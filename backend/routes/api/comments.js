/**
 * Express router providing comment related routes.
 * @module routes/api/comments
 */

const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Route to get all comments sorted by creation date in descending order.
 * @name get/comments
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {Promise<void>} - A promise that resolves to sending the list of comments
 */

/**
 * Route to delete a comment by ID.
 * @name delete/comments/:id
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object
 * @param {Object} res - Express response object
 * @returns {Promise<void>} - A promise that resolves to sending the deleted comment or an error message
 */

module.exports = router;

router.get("/comments", async (req, res) => {
  const comments = await Comment.find().sort({ createdAt: -1 });
  res.send(comments);
});

router.delete("/comments/:id", async (req, res) => {
    try {
        const comment = await Comment.findByIdAndDelete(req.params.id);
        if (!comment) {
            return res.status(404).send("Comment not found");
        }
        res.send(comment);
    } catch (error) {
        res.status(500).send(error.message);
    }
});
