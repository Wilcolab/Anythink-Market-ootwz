const router = require("express").Router();
const mongoose = require("mongoose");

const Comment = mongoose.model("Comment");

module.exports = router;

/**
 * GET all comments.
 * @route GET /api/comments
 * @returns {Object} - JSON object containing the comments.
 * @throws {Error} - If an error occurs while retrieving the comments.
 */

router.get("/", async (req, res) => {
    try {
        const comments = await Comment.find();
        res.json({ comments });
    } catch (err) {
        console.error(err);
    }
});

/**
 * DELETE a comment by ID.
 * @route DELETE /api/comments/:id
 * @param {string} id - The ID of the comment to be deleted.
 * @returns {Object} - JSON object indicating the success of the deletion.
 * @throws {Error} - If an error occurs while deleting the comment.
 */

router.delete("/:id", async (req, res) => {
    try {
        await Comment.findByIdAndRemove(req.params.id);
        res.json({ success: true });
    } catch (err) {
        console.error(err);
    }
});
