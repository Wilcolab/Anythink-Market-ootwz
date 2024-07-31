const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * Express router for handling comments API requests.
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
router.get("/", (req, res) => {
    Comment.find()
        .then(comments => {
            res.json(comments);
        })
        .catch(err => console.log(err));
});

/**
 * DELETE /api/comments/:id
 * Deletes a comment by its ID.
 * @name DELETE /api/comments/:id
 * @function
 * @memberof module:routes/api/comments
 * @param {Object} req - Express request object.
 * @param {Object} res - Express response object.
 * @returns {Object} - JSON response containing the deleted comment.
 */
router.delete("/:id", async (req, res) => {
        try {
                const comment = await Comment.findByIdAndRemove(req.params.id);
                res.json(comment);
        } catch (err) {
                console.log(err);
        }
});