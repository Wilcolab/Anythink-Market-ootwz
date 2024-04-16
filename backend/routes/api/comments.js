const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

/**
 * @route GET /api/comments
 * @desc Get all comments
 * @access Public
 */
router.get("/", (req, res, next) => {
    Comment.find()
        .then(comments => {
            return res.json({ comments: comments.map(comment => comment.toJSONFor()) });
        })
        .catch(next);
});

/**
 * @route DELETE /api/comments/:comment
 * @desc Delete a comment
 * @access Public
 * @param {string} req.params.comment - The ID of the comment to be deleted
 */
router.delete("/:comment", async (req, res, next) => {
    try {
        const comment = await Comment.findById(req.comment._id);
        if (!comment) {
            return res.sendStatus(404);
        }
        await comment.remove();
        return res.sendStatus(204);
    } catch (error) {
        next(error);
    }
});
