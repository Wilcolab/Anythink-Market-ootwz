const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

router.get("/", async (req, res) => {
    const comments = await Comment.find();
    res.json(comments);
    }
);

router.delete("/:id", async (req, res) => {
    try {
        /**
         * Represents a comment object.
         * @typedef {Object} Comment
         * @property {string} id - The unique identifier of the comment.
         * @property {string} content - The content of the comment.
         * @property {string} author - The author of the comment.
         * @property {Date} createdAt - The date and time when the comment was created.
         * @property {Date} updatedAt - The date and time when the comment was last updated.
         */
        const comment = await Comment.findById(req.params.id);
        if (!comment) {
            return res.status(404).json({ message: "Comment not found" });
        }
        await comment.remove();
        res.json({ message: "Comment removed" });
    } catch (error) {
        res.status(500).json({ message: "Internal server error" });
    }
});
