---- MODULE ProducerConsumer ----
EXTENDS Integers, Sequences

VARIABLES 
    ch,           \* 通道内容
    produced,     \* 已生产的消息数
    consumed,     \* 已消费的消息数
    closed        \* 通道是否关闭

TypeOK ==
    /\ ch \in Seq(0..4)
    /\ produced \in 0..5
    /\ consumed \in 0..5
    /\ closed \in BOOLEAN

Init ==
    /\ ch = <<>>
    /\ produced = 0
    /\ consumed = 0
    /\ closed = FALSE

Produce ==
    /\ produced < 5
    /\ ch = <<>>
    /\ ~closed
    /\ ch' = Append(ch, produced)
    /\ produced' = produced + 1
    /\ UNCHANGED <<consumed, closed>>

Close ==
    /\ produced = 5
    /\ ch = <<>>
    /\ ~closed
    /\ closed' = TRUE
    /\ UNCHANGED <<ch, produced, consumed>>

Consume ==
    /\ ch /= <<>>
    /\ ch' = Tail(ch)
    /\ consumed' = consumed + 1
    /\ UNCHANGED <<produced, closed>>

Next ==
    \/ Produce
    \/ Close
    \/ Consume

Fairness ==
    /\ SF_<<ch, produced, closed>>(Produce)
    /\ SF_<<produced, closed>>(Close)
    /\ SF_<<ch>>(Consume)

Spec == Init /\ [][Next]_<<ch, produced, consumed, closed>> /\ Fairness

THEOREM Spec => []TypeOK

ChannelEventuallyEmpty == <>(ch = <<>>)
AllMessagesProduced == <>(produced = 5)
ChannelEventuallyClosed == <>(closed = TRUE)
AllMessagesConsumed == <>(consumed = 5)

====
