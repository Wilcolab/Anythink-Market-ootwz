const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

/**
 * GET all comments
 * @route GET /api/comments
 * @returns {Array} Array of comments
 */
router.get("/", (req, res) => {
    Comment.find()
    .then((comments) => {
        res.json(comments);
    })
    .catch((err) => {
        res.status(500).json({ error: "Failed to retrieve comments" });
    });
});

/**
 * DELETE a comment by ID
 * @route DELETE /api/comments/:id
 * @param {string} req.params.id - The ID of the comment to delete
 * @returns {Object} Empty response with status 204 if successful, or status 500 if an error occurs
 */
router.delete("/:id", async (req, res) => {
    try {
        await Comment.findByIdAndDelete(req.params.id);
        res.status(204).send();
    } catch (err) {
        res.status(500).send();
    }
});

module.exports = router;
