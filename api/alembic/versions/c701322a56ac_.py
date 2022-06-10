"""Add layout metadata to topics

Revision ID: c701322a56ac
Revises: 635ef4648b17
Create Date: 2022-05-03 16:33:46.237136

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'c701322a56ac'
down_revision = '635ef4648b17'
branch_labels = None
depends_on = None

SQL = """
ALTER TABLE topic ADD COLUMN IF NOT EXISTS meta VARCHAR
"""

def upgrade():
    op.execute(SQL)



def downgrade():
    pass
