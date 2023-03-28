-- name: CreatePolls :exec
INSERT INTO polls (poll_by,poll_question,options_count,options)
VALUES (?,?,?,?);

-- name: GetPolls :many
select polls.poll_id,polls.poll_question,polls.options_count,polls_reaction.type,count(*) from polls
inner join polls_reaction on polls.poll_id = polls_reaction.poll_id
group by polls.poll_id, polls.poll_question, polls.options_count,polls_reaction.type