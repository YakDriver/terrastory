data "aws_iam_policy_document" "assume-role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "role" {
	name                = "${var.role_name}"
	assume_role_policy  = "${data.aws_iam_policy_document.assume-role.json}"
}

resource "aws_iam_policy" "managed-policy" {
  name = "${var.managed_policy_name}"
  path = "/tf-testing/"

  policy = <<EOF
{
	"Version": "2012-10-17",
	"Statement": [
		{
		"Action": [
			"ec2:Describe*"
		],
		"Effect": "Allow",
		"Resource": "*"
		}
	]
}
EOF
}

resource "aws_iam_policy_attachment" "test-attach" {
  name       = "test-attachment"
  roles      = ["${aws_iam_role.role.name}"]
  policy_arn = "${aws_iam_policy.managed-policy.arn}"
}
