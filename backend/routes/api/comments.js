const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

router.get("/", (req, res) => {
    Comment.find()
        .then((comments) => {
        res.json(comments);
        })
        .catch((err) => {
        console.log(err);
        });
    }
);
// copilot, implement a delete endpoint for comments
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
        const comment = await Comment.findByIdAndDelete(req.params.id);
        res.json(comment);
    } catch (err) {
        console.log(err);
    }
});
