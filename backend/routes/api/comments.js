/**
 * Express router for handling comments.
 * @type {import("express").Router}
 */
const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

/**
 * GET /api/comments
 * Retrieves all comments.
 * @param {import("express").Request} req - The request object.
 * @param {import("express").Response} res - The response object.
 */
router.get("/", (req, res) => {
    Comment.find()
            .then(comments => {
                res.json({ comments });
            })
            .catch(err => {
                console.log(err);
            });
});

/**
 * DELETE /api/comments/:id
 * Deletes a comment by ID.
 * @param {import("express").Request} req - The request object.
 * @param {import("express").Response} res - The response object.
 */
router.delete("/:id", async (req, res) => {
        try {
                await Comment.findByIdAndRemove(req.params.id);
                res.json({ success: true });
        } catch (err) {
                console.log(err);
        }
});
